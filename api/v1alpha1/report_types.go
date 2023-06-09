package v1alpha1

import (
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReportSpec defines the desired state of Report
type ReportSpec struct {
	//Namespace in which cron job is created
	CronjobNamespace string `json:"cronjobNamespace,omitempty"`

	//Schedule period for the Report Cronjob
	//ScheduleForReport string `json:"scheduleForReport,omitempty"`

	// Value for the Database Name Environment Variable
	DatabaseName string `json:"databaseName,omitempty"`

	//Value for the Database Password Environment Variable
	DatabasePassword string `json:"databasePassword,omitempty"`

	// Value for the Database User Environment Variable
	DatabaseUser string `json:"databaseUser,omitempty"`

	// Value for the Database HostName Environment Variable
	DatabaseHostName string `json:"databaseHostName,omitempty"`

	// Value for the Database Environment Variable in order to define the port which it should use. It will be used in its container as well
	DatabasePort string `json:"databasePort,omitempty"`

	// Frequency of the report generation
	ReportFrequency string `json:"reportFrequency,omitempty"`
}

// ReportStatus defines the observed state of Report
type ReportStatus struct {
	//Name of the CronJob object created and managed by it
	CronJobName string `json:"cronJobName"`

	//CronJobStatus represents the current state of a cronjob
	CronJobStatus batchv1.CronJobStatus `json:"cronJobStatus"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Report is the Schema for the reports API
type Report struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReportSpec   `json:"spec,omitempty"`
	Status ReportStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ReportList contains a list of Report
type ReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Report `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Report{}, &ReportList{})
}
