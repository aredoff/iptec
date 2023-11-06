package iptec

type HttpClient struct {
}

// func (c *HttpClient) MakeRequest[T any](req *fasthttp.Request) (*T, error) {
// 	mu.RLock()
// 	req.Header.Add("Authorization", "Bearer "+token)
// 	mu.RUnlock()
// 	req.Header.Add("User-Agent", "Reagate/1.0")
// 	attept := 0
// 	var err error
// 	resp := fasthttp.AcquireResponse()
// 	defer fasthttp.ReleaseResponse(resp)
// 	for attept < 3 {
// 		if attept > 0 {
// 			time.Sleep(time.Duration(attept) * time.Second)
// 		}
// 		attept++
// 		err = client.Do(req, resp)
// 		if resp.StatusCode() == 401 {
// 			err := Auth()
// 			if err != nil {
// 				return nil, err
// 			}
// 			mu.RLock()
// 			req.Header.Set("Authorization", "Bearer "+token)
// 			mu.RUnlock()
// 			continue
// 		}
// 		if err == nil {
// 			var responseData *T
// 			responseData, err = processingResponse[T](resp)
// 			if err != nil {
// 				continue
// 			} else {
// 				return responseData, nil
// 			}
// 		}
// 	}
// 	return nil, err
// }
