type Commit struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`
	Spec CommitSpec `json:"spec,omitempty"`
	Status CommitStatus `json:"status,omitempty"`
}
type CommitSpec struct {
	Images []Image `json:"images"`
	RestartPolicy RestartPolicy `json:"restartPolicy,omitempty"`
	PodName	string	`json:"podName,omitempty"`
	NodeName string `json:"nodeName,omitempty"`
	ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets,omitempty"`
}
type Image struct {
	ContainerName string `json:"containerName"`
	From string `json:"from"`
	Image string `json:"image"`
	ImageID string `json:"imageID"`
	ImagePushPolicy PushPolicy `json:"imagePushPolicy"`
}
type ContainerCommitStatus {
	ContainerName string `json:"containerName"`
	Phase                CommitState `json:"phase,omitempty"`
	RestartCount int `json:"restartCount"`
}
type CommitStatus struct {
	ContainerCommitStatuses []ContainerCommitStatus `json:"containerCommitStatus,omitempty"`
}