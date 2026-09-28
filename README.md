# Go + Test IT + GitHub Actions

Пример автотестов на Go, которые запускаются из Test IT через GitHub Actions и отправляют результаты обратно в Test IT с помощью [adapters-go](https://github.com/testit-tms/adapters-go).

**Проект в Test IT:** [team-0tm5.testit.software/projects/525/autotests](https://team-0tm5.testit.software/projects/525/autotests)

## Как это работает

1. Test IT отправляет webhook в GitHub — событие `repository_dispatch` с типом `run-tests`.
2. Запускается workflow [.github/workflows/.github-ci.yml](.github/workflows/.github-ci.yml).
3. В workflow устанавливается последняя версия адаптера (`go get github.com/testit-tms/adapters-go/v2@latest`).
4. Тесты выполняются через `go test`, результаты загружаются в Test IT.
5. К прогону в Test IT прикрепляется ссылка на пайплайн GitHub Actions (`actions/runs/<run_id>`).

### Режимы запуска

Режим задаётся полем `adapter_mode` в webhook.

| `adapter_mode` | Что происходит | Имя прогона |
|---|---|---|
| `0` | Результаты пишутся в существующий прогон, `test_run_id` берётся из webhook. Sync-storage запускается в workflow. | `GitHub Actions #<run_number> (adapterMode=0)` |
| `2` | Адаптер сам создаёт новый прогон, `test_run_id` не передаётся. | `GitHub Actions #<run_number> (adapterMode=2)` |

### Данные из webhook

```json
{
  "event_type": "run-tests",
  "client_payload": {
    "adapter_mode": "0",
    "url": "https://team-0tm5.testit.software",
    "project_id": "<id проекта>",
    "configuration_id": ["<id конфигурации>"],
    "test_run_id": "<id прогона, только для adapter_mode=0>"
  }
}
```

### Секреты репозитория

| Секрет | Назначение |
|---|---|
| `TMS_PRIVATE_TOKEN` | Приватный токен пользователя Test IT |

## Структура проекта

* **.github/workflows/.github-ci.yml** – workflow запуска тестов по webhook из Test IT
* **go.mod** – Go-модуль; адаптер добавляется в него в CI
* **examples/** – тесты
    * **main_test.go** – `TestMain` с `tms.Run(m)`, обязателен для адаптера v2
    * **before_after_test.go** – примеры setup/teardown
    * **metadata_test.go** – примеры [метаданных adapters-go (→ github.com)](https://github.com/testit-tms/adapters-go#usage)
    * **methods_test.go** – примеры [методов adapters-go (→ github.com)](https://github.com/testit-tms/adapters-go#usage): сообщения, ссылки, вложения
    * **parametrize_test.go** – примеры параметризованных тестов
    * **steps_test.go** – примеры шагов
    * **attachments/** – файлы для тестов с вложениями
