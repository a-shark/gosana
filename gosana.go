package gosana

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultBaseUrl = "https://app.asana.com/api/1.0"
)

type Projects struct {
	Data []struct { // How do I get rid off this struct "Data" ?
		Id   float64 `json:id`
		Name string  `json:name`
	}
}

type Project struct {
	Data struct {
		Id         float64   `json:id`
		Color      string    `json:color`
		CreatedAt  time.Time `json:created_at`
		ModifiedAt time.Time `json:modified_at`
		Archived   bool      `json:archived`
		Public     bool      `json:public`
		Name       string    `json:name`
		Notes      string    `json:name`

		Members []struct {
			Id   float64 `json:id`
			Name string  `json:name`
		}
		Followers []struct {
			Id   float64 `json:id`
			Name string  `json:name`
		}
		Workspace struct {
			Id   float64 `json:id`
			Name string  `json:name`
		}
	}
}

type Tasks struct {
	Data []struct {
		Id   float64 `json:id`
		Name string  `json:name`
	}
}

type Task struct {
	Data struct {
		Id        float64 `json:id`
		Name      string  `json:name`
		Notes     string  `json:notes`
		Completed bool    `json:completed`

		Tags []struct {
			Id   float64 `json:id`
			Name string  `json:name`
		}
	}
}

// Full task object WIP
// dunno if this much info is needed
//type Task struct {
//Data []struct {
//Id             float64   `json:id`
//Name           string    `json:name`
//Notes          string    `json:notes`
//AssigneeStatus string    `json:assignee_status`
//NumHearts      int       `json:num_hearts`
//Hearted        bool      `json:hearted`
//Completed      bool      `json:completed`
//DueOn          string    `json:due_on` // e.g: 2015-03-23
//CompletedAt    time.Time `json:completed_at`
//CreatedAt      time.Time `json:created_at`
//ModifiedAt     time.Time `json:modified_at`
//}
//}

// Client

type Client struct {
	User     string
	Password string // field used as default value, blank
}

func NewClient(password string) Client {
	return Client{User: password}
}

func (cli *Client) Request(endpoint string) ([]byte, error) {

	tr := &http.Transport{
		// Get rid off this bad habit, verify certs...
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}

	// Prepare request
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", defaultBaseUrl+endpoint, nil)
	req.Header.Add("User-Agent", "Gosana/0.1")
	req.SetBasicAuth(cli.User, cli.Password) // Asana sends <user>: so pwd is blank
	if err != nil {
		return nil, err
	}

	// Execute it
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Client - Asana endpoints

// Projects
func (cli *Client) Projects() Projects {
	body, _ := cli.Request("/projects")
	var projects Projects
	json.Unmarshal(body, &projects)
	return projects
}

func (cli *Client) Project(id string) {
	body, _ := cli.Request("/projects/" + id)
	var project Project
	json.Unmarshal(body, &project)
	fmt.Println(project)

}

// Tasks

func (cli *Client) Tasks(projectId string) Tasks {
	body, _ := cli.Request("/projects/" + projectId + "/tasks")
	var tasks Tasks
	json.Unmarshal(body, &tasks)
	return tasks
}

func (cli *Client) Task(taskId string) Task {
	body, _ := cli.Request("/tasks/" +
		taskId + "?opt_fields=name,notes,completed,tags.name") // few fields
	var task Task
	json.Unmarshal(body, &task)
	return task
}
