package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetUserBookList(login string) *ResponseBookSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/users/%s/repos", login)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func GetGroupBookList(token, groupLogin string) *ResponseBookSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s/repos", groupLogin)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func (c *Client) GetBookInfo(namespace string) *ResponseBookDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s", namespace)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func GetBookInfo(token, namespace string) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s", namespace)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func (c *Client) CreateUserBook(login, changeSlug, changeName string) *ResponseBookDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/users/%s/repos?slug=%s&name=%s", login, changeSlug, changeName)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func CreateUserBook(token, login, changeSlug, changeName string) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/users/%s/repos?slug=%s&name=%s", login, changeSlug, changeName)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func (c *Client) CreateGrouprBook(login, changeSlug, changeName string, changePublic int) *ResponseBookDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s/repos?slug=%s&name=%s&public=%d", login, changeSlug, changeName, changePublic)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

/*public:
0  private
1  public
2  visible for space members
3  visible for everyone int space members
4  visible for book members
*/
func CreateGrouprBook(token, login, changeSlug, changeName string, changePublic int) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/groups/%s/repos?slug=%s&name=%s&public=%d", login, changeSlug, changeName, changePublic)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func (c *Client) UpdateBook(namespace, changeName, changeSlug string) *ResponseBookDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s?name=%s&slug=%s", namespace, changeName, changeSlug)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func UpdateBook(token, namespace, changeName, changeSlug string) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s?name=%s&slug=%s", namespace, changeName, changeSlug)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func (c *Client) DeleteBook(namespace string) *ResponseBookDetailSerializer {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s", namespace)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func DeleteBook(token, namespace string) *ResponseBookDetailSerializer {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s", namespace)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDetailSerializer{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

func (c *Client) GetBookDirectoryStructure(namespace string) *ResponseBookDirectoryStructure {
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/toc", namespace)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDirectoryStructure{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}

//need repo abilities.read authority
func GetBookDirectoryStructure(token, namespace string) *ResponseBookDirectoryStructure {
	client := http.DefaultClient
	url := fmt.Sprintf("https://www.yuque.com/api/v2/repos/%s/toc", namespace)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	book := ResponseBookDirectoryStructure{}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err)
	}

	return &book
}
