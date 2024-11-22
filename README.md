# Image Resampling Service

Сервис принимает изображение в формате JPEG в кодировке base64, преобразует его в заданный размер и сохраняет обработанное изображение на диске. Он также кэширует обработанные изображения, чтобы избежать лишних операций.


## Начало работы

### Требования

**Docker**: Убедитесь, что у вас установлен Docker. Скачать его можно с [сайта Docker](https://www.docker.com/products/docker-desktop/).
**Go 1.21+**: Для локального запуска


### Установка и конфигурация

1. **Флаги**

    ```plaintext
    -path-orig <str>: Директория для исходного изображения. По умолчанию: /tmp/img_orig.
    -path-res <str>: Директория для обработанных изображений. По умолчанию: /tmp/img_res.
    -width <uint>: Ширина обработанного изображения. По умолчанию: 200.
    -height <uint>: Длина обработанного изображения. По умолчанию: 200.
    ```

2. **Старт**

    ```
    docker build -t img-resampler .
    docker run -p 8085:8085 img-resampler
    ```

3. **Для локального запуска**
    ```
    go mod tidy
    go build -o img-resampler ./cmd
    ./img-resampler
    ```

## API Endpoints

### POST /process

**Тело запроса**
```
{
  "image": "<base64-encoded-image>"
}
```

**Ответ**
```
{
  "time": 123,         // Время обработки в миллисекундах
  "cached": true|false // Использовался ли кэш
}
```

**Пример запроса (где output.txt содержит изображение в кодировке base64)**:

```curl -X POST http://localhost:8085/process -H "Content-Type: application/json" -d "{\"image\":\"$(cat output.txt)\"}"```