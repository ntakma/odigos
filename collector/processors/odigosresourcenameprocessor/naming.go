package odigosresourcenameprocessor

import (
	"context"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ContainerDetails struct {
	PodName         string
	PodNamespace    string
	ContainersInPod int
	ContainerName   string
}

type K8sResourceAttributes struct {
	OtelServiceName string
	Namespace       string
	WorkloadKind    string
	WorkloadName    string
	PodName         string
	ContainerName   string
}

type NameStrategy interface {
	GetK8sResourceAttributes(containerDetails *ContainerDetails) *K8sResourceAttributes
}

type NameFromOwner struct {
	kc     kubernetes.Interface
	logger *zap.Logger
}

func (n *NameFromOwner) GetK8sResourceAttributes(containerDetails *ContainerDetails) *K8sResourceAttributes {

	attrs := K8sResourceAttributes{
		Namespace:     containerDetails.PodNamespace,
		PodName:       containerDetails.PodName,
		ContainerName: containerDetails.ContainerName,
	}

	if containerDetails.ContainersInPod > 1 {
		attrs.OtelServiceName = containerDetails.ContainerName
		return &attrs
	}

	workloadName, workloadKind, err := n.getWorkloadNameByOwner(containerDetails)
	if err != nil {
		n.logger.Error("Failed to get name by owner, using pod name", zap.Error(err))
		attrs.OtelServiceName = containerDetails.PodName
		return &attrs
	}

	attrs.WorkloadKind = workloadKind
	attrs.WorkloadName = workloadName

	overwrittenName, exists := n.getServiceNameFromAnnotation(workloadName, workloadKind, containerDetails.PodNamespace)
	if !exists {
		attrs.OtelServiceName = workloadName
		return &attrs
	}

	attrs.OtelServiceName = overwrittenName
	return &attrs
}

func (n *NameFromOwner) getServiceNameFromAnnotation(name string, kind string, namespace string) (string, bool) {
	obj := n.getKubeObject(name, kind, namespace)
	if obj == nil {
		return "", false
	}

	annotations := obj.GetAnnotations()
	if annotations == nil {
		return "", false
	}

	overwrittenName, exists := annotations["odigos.io/reported-name"]
	if !exists {
		return "", false
	}

	return overwrittenName, true
}

func (n *NameFromOwner) getKubeObject(name string, kind string, namespace string) metav1.Object {
	switch kind {
	case "Deployment":
		deployment, err := n.kc.AppsV1().Deployments(namespace).Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			n.logger.Error("Failed to get deployment", zap.Error(err))
			return nil
		}

		return deployment
	case "StatefulSet":
		statefulSet, err := n.kc.AppsV1().StatefulSets(namespace).Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			n.logger.Error("Failed to get statefulset", zap.Error(err))
			return nil
		}

		return statefulSet
	case "DaemonSet":
		daemonSet, err := n.kc.AppsV1().DaemonSets(namespace).Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			n.logger.Error("Failed to get daemonset", zap.Error(err))
			return nil
		}

		return daemonSet
	case "Pod":
		pod, err := n.kc.CoreV1().Pods(namespace).Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			n.logger.Error("Failed to get pod", zap.Error(err))
			return nil
		}

		return pod
	}

	return nil
}

func (n *NameFromOwner) getWorkloadNameByOwner(containerDetails *ContainerDetails) (workloadName string, workloadKind string, err error) {
	pod, err := n.kc.CoreV1().Pods(containerDetails.PodNamespace).
		Get(context.Background(), containerDetails.PodName, metav1.GetOptions{})
	if err != nil {
		return "", "", err
	}

	ownerRefs := pod.GetOwnerReferences()
	for _, ownerRef := range ownerRefs {
		if ownerRef.Kind == "ReplicaSet" {
			rs, err := n.kc.AppsV1().ReplicaSets(pod.Namespace).Get(context.Background(), ownerRef.Name, metav1.GetOptions{})
			if err != nil {
				return "", "", err
			}

			ownerRefs = rs.GetOwnerReferences()
			for _, ownerRef := range ownerRefs {
				if ownerRef.Kind == "Deployment" {
					return ownerRef.Name, ownerRef.Kind, nil
				}
			}
		} else if ownerRef.Kind == "StatefulSet" || ownerRef.Kind == "DaemonSet" {
			return ownerRef.Name, ownerRef.Kind, nil
		}
	}

	return containerDetails.PodName, "Pod", nil
}
