package storage

//type QiniuStorage struct {
//	AccessKey string
//	SecretKey string
//	Bucket    string
//	Domain    string
//}
//
//var _ Storage = (*QiniuStorage)(nil) // 确保实现了Storage接口
//
//func (q *QiniuStorage) GetUrl(key string) (string, error) {
//	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
//	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
//	privateAccessURL := storage.MakePrivateURL(mac, q.Domain, key, deadline)
//	return privateAccessURL, nil
//}
