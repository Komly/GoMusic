package vkapi

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

const responseMock = `
{"response":{"count":86,"items":[{"id":340388,"from_id":-1,"owner_id":-1,"date":1454599127,"post_type":"post","text":"Долгожданная возможность получать материалы из диалогов и групповых чатов теперь реализована в API — добавлен метод vk.com\/dev\/messages.getHistoryAttachments\n\nС его помощью Вы можете работать с фотографиями, видео, аудиозаписями, документами и внешними ссылками, которые были отправлены или получены пользователем в рамках диалога или чата, без необходимости фильтровать сообщения по наличию вложений на своей стороне.\n\nЗа один вызов метод возвращает до 200 объектов выбранного типа. Для подгрузки большего числа материалов используйте параметр start_from, новое значение для которого возвращается вместе с ответом в поле next_from.","comments":{"count":0},"likes":{"count":220},"reposts":{"count":18}}]}}`

func TestGetURL(t *testing.T) {
	apiClient := NewAPIClient()
	resultURL := apiClient.getRequestURL("wall.get", nil)

	if resultURL.String() != "https://api.vk.com/method/wall.get?v=5.37" {
		t.Error("Invalid url", resultURL.String())
	}

	params := map[string]string{"owner_id": "-1"}

	result := apiClient.getRequestURL("wall.get", params)
	if result.String() != "https://api.vk.com/method/wall.get?owner_id=-1&v=5.37" {
		t.Error("invalir url", result)
	}
}

func TestGetRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.vk.com/method/wall.get?count=1&owner_id=-1&v=5.37",
		httpmock.NewStringResponder(200, responseMock))

	apiClient := NewAPIClient()
	params := map[string]string{"owner_id": "-1", "count": "1"}
	_, err := apiClient.APIRequest("wall.get", params)
	if err != nil {
		t.Error(err)
	}

}

func TestFailedRequest(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "https://api.vk.com/method/wall.get?count=1&owner_id=-1&v=5.37",
		func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, errors.New("Aww")
		})
	defer httpmock.DeactivateAndReset()

	apiClient := NewAPIClient()
	params := map[string]string{"owner_id": "-1", "count": "1"}
	_, err := apiClient.APIRequest("wall.get", params)

	if err == nil {
		t.Error("Expected to get error")
	}

}
