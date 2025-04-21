# TF-IDF Web Application

Веб-приложение для расчета метрик TF-IDF из текстовых файлов.

## Запуск
```bash
go run main.go
```

## API
- **Метод**: POST
- **URL**: localhost:8080/tfidf
- **Content-Type**: multipart/form-data
- **Параметр формы**: txtFile (текстовый файл)

## Результат
- JSON с 50 словами, упорядоченными по убыванию IDF
- Для каждого слова выводится: слово, TF, IDF, TF-IDF
