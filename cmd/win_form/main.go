package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	_ "github.com/lib/pq"
)

/*
	func getTableData(db *sql.DB) ([][]string, error) {
		// Выполнение запроса к базе данных для получения данных
		rows, err := db.Query("SELECT name, email FROM admins")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// Инициализация двумерного слайса для хранения данных таблицы
		var data [][]string

		// Чтение данных из результатов запроса и добавление их в слайс
		for rows.Next() {
			var name, email string
			err := rows.Scan(&name, &email)
			if err != nil {
				return nil, err
			}
			data = append(data, []string{name, email})
		}

		return data, nil
	}
*/
func main() {
	// Подключение к базе данных PostgreSQL
	//dsn := "postgres://user:password@localhost:5432/database_name?sslmode=disable"
	//db, err := sql.Open("postgres", dsn)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()

	// Создание экземпляра приложения Fyne
	myApp := app.New()

	// Создание основного контейнера приложения
	win := myApp.NewWindow("Пример формы с Fyne")

	// Создание кнопок
	button1 := widget.NewButton("Кнопка 1", func() {
		// Обработчик нажатия на кнопку 1
	})
	button2 := widget.NewButton("Кнопка 2", func() {
		// Обработчик нажатия на кнопку 2
	})
	button3 := widget.NewButton("Кнопка 3", func() {
		// Обработчик нажатия на кнопку 3
	})
	button4 := widget.NewButton("Кнопка 4", func() {
		// Обработчик нажатия на кнопку 4
	})

	// Создание текстовых полей
	textField1 := widget.NewEntry()
	textField2 := widget.NewEntry()
	textField3 := widget.NewEntry()
	textField4 := widget.NewEntry()

	// Создание виджета таблицы
	//table := widget.NewTable(
	//	func() (int, int) {
	//		// Получение данных из базы данных для определения количества строк и столбцов
	//		data, err := getTableData(db)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		// Возвращение количества строк и столбцов
	//		return len(data), len(data[0])
	//	},
	//	func() fyne.CanvasObject {
	//		// Функция, возвращающая содержимое каждой ячейки таблицы
	//		// В данном примере таблица заполняется данными из базы данных
	//
	//		// Создание пустого текстового виджета
	//		label := widget.NewLabel("")
	//		return label
	//	},
	//)

	// Заполнение таблицы данными из базы данных
	//data, err := getTableData(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for row, rowData := range data {
	//	for col, cellData := range rowData {
	//		label := widget.NewLabel(cellData)
	//		table.SetCell(table.NewCell(label, row, col))
	//	}
	//}

	// Создание вертикального контейнера для кнопок
	buttonsContainer := container.NewVBox(button1, button2, button3, button4)

	// Создание горизонтального контейнера для текстовых полей
	textFieldsContainer := container.NewHBox(textField1, textField2, textField3, textField4)

	// Создание вертикального контейнера для всех элементов формы
	formContainer := container.NewVBox(
		buttonsContainer,
		textFieldsContainer,
		//table,
	)

	// Добавление контейнера с элементами формы на окно приложения
	win.SetContent(formContainer)

	// Отображение окна приложения
	win.ShowAndRun()
}
