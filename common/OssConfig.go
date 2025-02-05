package common

type OSSConfigModule struct {
	EndPoint        string            `json:"endPoint"`
	AccessKeyID     string            `json:"accessKeyID"`
	AccessKeySecret string            `json:"accessKeySecret"`
	DefaultBucket   string            `json:"defaultBucket"`
	RoleArn         string            `json:"roleArn"`
	STSEndPoint     string            `json:"stsEndPoint"`
	Folder          map[string]string `json:"folder"`
}
