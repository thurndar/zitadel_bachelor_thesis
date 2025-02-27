package s3

/* Deprecated in V2

import (
	kubernetesmock "github.com/caos/orbos/pkg/kubernetes/mock"
	"github.com/caos/zitadel/operator/database/kinds/backups/bucket/backup"
	"github.com/caos/zitadel/operator/database/kinds/backups/bucket/restore"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	macherrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func SetInstantBackup(
	k8sClient *kubernetesmock.MockClientInt,
	namespace string,
	backupName string,
	labelsAKID map[string]string,
	labelsSAK map[string]string,
	labelsST map[string]string,
	akid, sak, st string,
) {
	k8sClient.EXPECT().ApplySecret(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      accessKeyIDName,
			Namespace: namespace,
			Labels:    labelsAKID,
		},
		StringData: map[string]string{accessKeyIDKey: akid},
		Type:       "Opaque",
	}).MinTimes(1).MaxTimes(1).Return(nil)

	k8sClient.EXPECT().ApplySecret(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretAccessKeyName,
			Namespace: namespace,
			Labels:    labelsSAK,
		},
		StringData: map[string]string{secretAccessKeyKey: sak},
		Type:       "Opaque",
	}).MinTimes(1).MaxTimes(1).Return(nil)

	k8sClient.EXPECT().ApplySecret(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sessionTokenName,
			Namespace: namespace,
			Labels:    labelsST,
		},
		StringData: map[string]string{sessionTokenKey: st},
		Type:       "Opaque",
	}).MinTimes(1).MaxTimes(1).Return(nil)

	k8sClient.EXPECT().ApplyJob(gomock.Any()).Times(1).Return(nil)
	k8sClient.EXPECT().GetJob(namespace, backup.GetJobName(backupName)).Times(1).Return(nil, macherrs.NewNotFound(schema.GroupResource{"batch", "jobs"}, backup.GetJobName(backupName)))
	k8sClient.EXPECT().WaitUntilJobCompleted(namespace, backup.GetJobName(backupName), gomock.Any()).Times(1).Return(nil)
	k8sClient.EXPECT().DeleteJob(namespace, backup.GetJobName(backupName)).Times(1).Return(nil)
}

func SetBackup(
	k8sClient *kubernetesmock.MockClientInt,
	namespace string,
	labelsAKID map[string]string,
	labelsSAK map[string]string,
	labelsST map[string]string,
	akid, sak, st string,
) {
	k8sClient.EXPECT().ApplySecret(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      accessKeyIDName,
			Namespace: namespace,
			Labels:    labelsAKID,
		},
		StringData: map[string]string{accessKeyIDKey: akid},
		Type:       "Opaque",
	}).MinTimes(1).MaxTimes(1).Return(nil)

	k8sClient.EXPECT().ApplySecret(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretAccessKeyName,
			Namespace: namespace,
			Labels:    labelsSAK,
		},
		StringData: map[string]string{secretAccessKeyKey: sak},
		Type:       "Opaque",
	}).MinTimes(1).MaxTimes(1).Return(nil)

	k8sClient.EXPECT().ApplySecret(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sessionTokenName,
			Namespace: namespace,
			Labels:    labelsST,
		},
		StringData: map[string]string{sessionTokenKey: st},
		Type:       "Opaque",
	}).MinTimes(1).MaxTimes(1).Return(nil)
	k8sClient.EXPECT().ApplyCronJob(gomock.Any()).Times(1).Return(nil)
}

func SetRestore(
	k8sClient *kubernetesmock.MockClientInt,
	namespace string,
	backupName string,
	labelsAKID map[string]string,
	labelsSAK map[string]string,
	labelsST map[string]string,
	akid, sak, st string,
) {
	k8sClient.EXPECT().ApplySecret(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      accessKeyIDName,
			Namespace: namespace,
			Labels:    labelsAKID,
		},
		StringData: map[string]string{accessKeyIDKey: akid},
		Type:       "Opaque",
	}).MinTimes(1).MaxTimes(1).Return(nil)

	k8sClient.EXPECT().ApplySecret(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretAccessKeyName,
			Namespace: namespace,
			Labels:    labelsSAK,
		},
		StringData: map[string]string{secretAccessKeyKey: sak},
		Type:       "Opaque",
	}).MinTimes(1).MaxTimes(1).Return(nil)

	k8sClient.EXPECT().ApplySecret(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sessionTokenName,
			Namespace: namespace,
			Labels:    labelsST,
		},
		StringData: map[string]string{sessionTokenKey: st},
		Type:       "Opaque",
	}).MinTimes(1).MaxTimes(1).Return(nil)
	k8sClient.EXPECT().ApplyJob(gomock.Any()).Times(1).Return(nil)
	k8sClient.EXPECT().GetJob(namespace, restore.GetJobName(backupName)).Times(1).Return(nil, macherrs.NewNotFound(schema.GroupResource{"batch", "jobs"}, restore.GetJobName(backupName)))
	k8sClient.EXPECT().WaitUntilJobCompleted(namespace, restore.GetJobName(backupName), gomock.Any()).Times(1).Return(nil)
	k8sClient.EXPECT().DeleteJob(namespace, restore.GetJobName(backupName)).Times(1).Return(nil)
}

*/
