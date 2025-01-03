/*
Copyright 2022.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package actions

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	actionv1 "github.com/odigos-io/odigos/api/actions/v1alpha1"
	v1 "github.com/odigos-io/odigos/api/odigos/v1alpha1"
	"github.com/odigos-io/odigos/common"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var supportedSignals = map[common.ObservabilitySignal]struct{}{
	common.TracesObservabilitySignal: {},
}

type ProbabilisticSamplerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *ProbabilisticSamplerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	logger := log.FromContext(ctx)
	logger.V(0).Info("Reconciling ProbabilisticSampler action")

	action := &actionv1.ProbabilisticSampler{}
	err := r.Get(ctx, req.NamespacedName, action)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	for _, signal := range action.Spec.Signals {

		if _, ok := supportedSignals[signal]; !ok {

			err = fmt.Errorf("unsupported signal: %s", signal)
			logger.V(0).Error(err, err.Error())
			r.ReportReconciledToProcessorFailed(ctx, action, FailedToTransformToProcessorReason, err.Error())
			return ctrl.Result{}, nil
		}
	}

	processor, err := r.convertToProcessor(action)
	if err != nil {
		logger.V(0).Error(err, "failed to convert ProbabilisticSampler to processor")
		r.ReportReconciledToProcessorFailed(ctx, action, FailedToTransformToProcessorReason, err.Error())
		return ctrl.Result{}, nil
	}

	err = r.Patch(ctx, processor, client.Apply, client.FieldOwner(action.Name), client.ForceOwnership)
	if err != nil {
		r.ReportReconciledToProcessorFailed(ctx, action, FailedToCreateProcessorReason, err.Error())
		return ctrl.Result{}, err
	}

	err = r.ReportReconciledToProcessor(ctx, action)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *ProbabilisticSamplerReconciler) ReportReconciledToProcessorFailed(ctx context.Context, action *actionv1.ProbabilisticSampler, reason string, msg string) error {
	changed := meta.SetStatusCondition(&action.Status.Conditions, metav1.Condition{
		Type:               ActionTransformedToProcessorType,
		Status:             metav1.ConditionFalse,
		Reason:             reason,
		Message:            msg,
		ObservedGeneration: action.Generation,
	})

	if changed {
		err := r.Status().Update(ctx, action)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ProbabilisticSamplerReconciler) ReportReconciledToProcessor(ctx context.Context, action *actionv1.ProbabilisticSampler) error {
	changed := meta.SetStatusCondition(&action.Status.Conditions, metav1.Condition{
		Type:               ActionTransformedToProcessorType,
		Status:             metav1.ConditionTrue,
		Reason:             ProcessorCreatedReason,
		Message:            "The action has been reconciled to a processor resource.",
		ObservedGeneration: action.Generation,
	})

	if changed {
		err := r.Status().Update(ctx, action)
		if err != nil {
			return err
		}
	}
	return nil
}

type ProbabilisticSamplerConfig struct {
	Value    float64 `json:"sampling_percentage"`
	HashSeed int     `json:"hash_seed"`
}

func (r *ProbabilisticSamplerReconciler) convertToProcessor(action *actionv1.ProbabilisticSampler) (*v1.Processor, error) {

	samplingPercentage, err := strconv.ParseFloat(action.Spec.SamplingPercentage, 32)
	if err != nil {
		return nil, err
	}

	if samplingPercentage < 0 {
		return nil, errors.New("sampling_precentage cannot be negative")
	}

	config := ProbabilisticSamplerConfig{
		Value: samplingPercentage,
		// Arbitrary hash seed set to maximize processor creating abstraction for the user.
		HashSeed: 123,
	}

	configJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	processor := v1.Processor{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "odigos.io/v1alpha1",
			Kind:       "Processor",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      action.Name,
			Namespace: action.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: action.APIVersion,
					Kind:       action.Kind,
					Name:       action.Name,
					UID:        action.UID,
				},
			},
		},
		Spec: v1.ProcessorSpec{
			Type:            "probabilistic_sampler",
			ProcessorName:   action.Spec.ActionName,
			Disabled:        action.Spec.Disabled,
			Notes:           action.Spec.Notes,
			Signals:         action.Spec.Signals,
			CollectorRoles:  []v1.CollectorsGroupRole{v1.CollectorsGroupRoleNodeCollector},
			OrderHint:       1,
			ProcessorConfig: runtime.RawExtension{Raw: configJson},
		},
	}

	return &processor, nil
}
