package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Name          string `json:"name"`
	Auths         string `json:"auths"`
	PubHouse      string `json:"pub_house"`
	PubYear       string `json:"pub_year"`
	CountInHall1  uint16 `json:"count_in_hall1"`
	CountInHall2  uint16 `json:"count_in_hall2"`
	CountInHall3  uint16 `json:"count_in_hall3"`
	Id            string `json:"book_id"`
	CountInHall1Z uint16 `json:"count_hall_1_z"`
	CountInHall2Z uint16 `json:"count_hall_2_z"`
	CountInHall3Z uint16 `json:"count_hall_3_z"`
}

type UserBooks struct {
	book1       string
	book2       string
	book3       string
	date_book_1 string
	date_book_2 string
	date_book_3 string
	hall        uint16
}

func returnBook(user_id, a uint16) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Подключено")

	defer db.Close()

	select_user_var, err := db.Query(fmt.Sprintf("SELECT `hall` FROM `USER` WHERE `USER`.`id` = %d", user_id)) // достаем номер зала, в который записан читатель
	if err != nil {
		panic(err)
	}

	var ifhall uint16 // 	инициализируем переменную, в которой будет записан этот номер

	for select_user_var.Next() {
		var user UserBooks
		err = select_user_var.Scan(&user.hall)
		if err != nil {
			panic(err)
		}

		ifhall = user.hall // записываем номер зала в переменную
	}

	select_user_var, err = db.Query(fmt.Sprintf("SELECT `book_%d` FROM `USER` WHERE `USER`.`id` = %d", a, user_id)) // достаем номер зала, в который записан читатель
	if err != nil {
		panic(err)
	}

	var book_id string // 	инициализируем переменную, в которой будет записан этот номер

	for select_user_var.Next() {
		var book Book
		err = select_user_var.Scan(&book.Id)
		if err != nil {
			panic(err)
		}

		book_id = book.Id // записываем номер зала в переменную
	}

	defer select_user_var.Close()

	select_book_var, err := db.Query(fmt.Sprintf("SELECT `count_hall_%d_z` FROM `BOOK` WHERE `BOOK`.`id` = '%s'", ifhall, book_id))
	if err != nil {
		panic(err)
	}

	var bc uint16

	for select_book_var.Next() {
		var book Book
		err = select_book_var.Scan(&book.CountInHall1Z)
		if err != nil {
			panic(err)
		}
		bc = book.CountInHall1Z
		fmt.Println(bc)

		switch a {
		case 1:
			select_var, err := db.Query(fmt.Sprintf("UPDATE `USER` SET `book_1` = NULL WHERE `USER`.`id` = %d", user_id))
			if err != nil {
				panic(err)
			}
			select_var_book, err := db.Query(fmt.Sprintf("UPDATE `BOOK` SET `count_hall_%d_z` = %d WHERE `BOOK`.`id` = '%s'", ifhall, bc-1, book_id))
			if err != nil {
				panic(err)
			}
			defer select_var.Close()
			defer select_var_book.Close()
		case 2:
			select_var, err := db.Query(fmt.Sprintf("UPDATE `USER` SET `book_2` = NULL WHERE `USER`.`id` = %d", user_id))
			if err != nil {
				panic(err)
			}
			select_var_book, err := db.Query(fmt.Sprintf("UPDATE `BOOK` SET `count_hall_%d_z` = %d WHERE `BOOK`.`id` = '%s'", ifhall, bc-1, book_id))
			if err != nil {
				panic(err)
			}
			defer select_var.Close()
			defer select_var_book.Close()
		case 3:
			select_var, err := db.Query(fmt.Sprintf("UPDATE `USER` SET `book_3` = NULL WHERE `USER`.`id` = %d", user_id))
			if err != nil {
				panic(err)
			}

			select_var_book, err := db.Query(fmt.Sprintf("UPDATE `BOOK` SET `count_hall_%d_z` = %d WHERE `BOOK`.`id` = '%s'", ifhall, bc-1, book_id))
			if err != nil {
				panic(err)
			}
			defer select_var.Close()
			defer select_var_book.Close()
		}

	}
}

func replaceBookInUsers(book_id string, id uint16) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Подключено")

	defer db.Close()

	select_user_var, err := db.Query(fmt.Sprintf("SELECT `hall` FROM `USER` WHERE `USER`.`id` = %d", id)) // достаем номер зала, в который записан читатель
	if err != nil {
		panic(err)
	}

	var ifhall uint16 // 	инициализируем переменную, в которой будет записан этот номер

	for select_user_var.Next() {
		var user UserBooks
		err = select_user_var.Scan(&user.hall)
		if err != nil {
			panic(err)
		}

		ifhall = user.hall // записываем номер зала в переменную
	}

	defer select_user_var.Close()

	select_book_var, err := db.Query(fmt.Sprintf("SELECT `count_in_hall%d`, `count_hall_%d_z` FROM `BOOK` WHERE `BOOK`.`id` = '%s'", ifhall, ifhall, book_id)) // достаем из значений сколько книг уже взято и сколько всего было
	if err != nil {
		panic(err)
	}

	var bc uint16

	for select_book_var.Next() {
		var book Book
		err = select_book_var.Scan(&book.CountInHall1, &book.CountInHall1Z)
		if err != nil {
			panic(err)
		}
		fmt.Println(book.CountInHall1Z)
		bc = book.CountInHall1Z
		fmt.Println(bc)
		fmt.Println(book.CountInHall1Z)

		if book.CountInHall1 <= bc { // проверяем, если количество взятых книг = количеству изначальных, то функция тут и заканчивается
			fmt.Println("В зале не осталось нужных книг")
		} else {

			select_var, err := db.Query(fmt.Sprintf("SELECT `book_1` FROM `USER` WHERE `USER`.`id` = %d", id))
			if err != nil {
				panic(err)
			}

			for select_var.Next() {
				var book Book
				err = select_var.Scan(&book.Name)
				if err == nil {
					select_var, err = db.Query(fmt.Sprintf("SELECT `book_2` FROM `USER` WHERE `USER`.`id` = %d", id))
					if err != nil {
						panic(err)
					}

					for select_var.Next() {
						var book Book
						err = select_var.Scan(&book.Name)
						if err == nil {
							select_var, err = db.Query(fmt.Sprintf("SELECT `book_3` FROM `USER` WHERE `USER`.`id` = %d", id))
							if err != nil {
								panic(err)
							}

							for select_var.Next() {
								var book Book
								err = select_var.Scan(&book.Name)
								if err == nil {
									fmt.Println("Читатель взял максимальное количество книг (3)")
								} else {
									select_var, err = db.Query(fmt.Sprintf("UPDATE `USER` SET `book_3` = '%s' WHERE `USER`.`id` = %d", book_id, id))
									if err != nil {
										panic(err)
									}

									select_var_book, err := db.Query(fmt.Sprintf("UPDATE `BOOK` SET `count_hall_%d_z` = %d WHERE `BOOK`.`id` = '%s'", ifhall, bc+1, book_id))
									if err != nil {
										panic(err)
									}
									defer select_var_book.Close()
								}
							}
						} else {
							select_var, err = db.Query(fmt.Sprintf("UPDATE `USER` SET `book_2` = '%s' WHERE `USER`.`id` = %d", book_id, id))
							if err != nil {
								panic(err)
							}

							select_var_book, err := db.Query(fmt.Sprintf("UPDATE `BOOK` SET `count_hall_%d_z` = %d WHERE `BOOK`.`id` = '%s'", ifhall, bc+1, book_id))
							if err != nil {
								panic(err)
							}
							fmt.Println(ifhall, bc+1)
							defer select_var_book.Close()
						}
					}
				} else {
					select_var, err = db.Query(fmt.Sprintf("UPDATE `USER` SET `book_1` = '%s' WHERE `USER`.`id` = %d", book_id, id))
					if err != nil {
						panic(err)
					}

					select_var_book, err := db.Query(fmt.Sprintf("UPDATE `BOOK` SET `count_hall_%d_z` = %d WHERE `BOOK`.`id` = '%s'", ifhall, bc+1, book_id))
					if err != nil {
						panic(err)
					}
					defer select_var_book.Close()
				}
			}
			defer select_var.Close()

		}
	}

	defer select_book_var.Close()

}

func selectBooksID(name string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Подключено")

	defer db.Close()

	select_var, err := db.Query(fmt.Sprintf("SELECT `id` FROM `BOOK` WHERE `BOOK`.`name` = '%s'", name))
	if err != nil {
		panic(err)
	}

	for select_var.Next() {
		var book Book
		err = select_var.Scan(&book.Id)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("Шифр книги: %s", book.Id))
	}

	defer select_var.Close()

}

func selectBooksName(id string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Подключено")

	defer db.Close()

	select_var, err := db.Query(fmt.Sprintf("SELECT `name` FROM `BOOK` WHERE `BOOK`.`id` = '%s'", id))
	if err != nil {
		panic(err)
	}

	for select_var.Next() {
		var book Book
		err = select_var.Scan(&book.Name)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("Название книги: %s", book.Name))
	}

	defer select_var.Close()

}

func selectUersBooks(id uint16) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Подключено")

	defer db.Close()

	select_var, err := db.Query(fmt.Sprintf("SELECT `book_1`, `date_book_1`, `book_2`, `date_book_2`, `book_3`, `date_book_3` FROM `USER` WHERE `USER`.`id` = %d", id))
	if err != nil {
		panic(err)
	}

	for select_var.Next() {
		var user UserBooks
		err = select_var.Scan(&user.book1, &user.date_book_1, &user.book2, &user.date_book_2, &user.book3, &user.date_book_3)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("У читателя с билетом %d взята книга %s в %s, книга %s в %s, книга %s в %s", id, user.book1, user.date_book_1, user.book2, user.date_book_2, user.book3, user.date_book_3))
	}

	defer select_var.Close()

}

func insertUser(last_name, date_birth, adress, academic_degree, admission_library string, passport_number, phone_number, hall uint32) {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Добавлено")

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `USER` (`last_name`, `passport_number`, `date_birth`, `adress`, `phone_number`, `academic_degree`, `hall`, `admission_library`) VALUES('%s', %d, '%s', '%s', %d, '%s', %d, '%s')", last_name, passport_number, date_birth, adress, phone_number, academic_degree, hall, admission_library)) //'2005-08-25'
	if err != nil {
		panic(err)
	}

	defer insert.Close()

	select_var, err := db.Query(fmt.Sprintf("SELECT `capacity` FROM `HALL` WHERE `HALL`.`number` = %d", hall))
	if err != nil {
		panic(err)
	}

	var current_count int

	for select_var.Next() {
		err = select_var.Scan(&current_count)
		if err != nil {
			panic(err)
		}

		fmt.Println(current_count)
	}

	fmt.Println(current_count)

	defer select_var.Close()

	minus_var, err := db.Query(fmt.Sprintf("UPDATE `HALL` SET `capacity` = %d WHERE `HALL`.`number` = %d", current_count-1, hall)) //'2005-08-25'
	if err != nil {
		panic(err)
	}

	defer minus_var.Close()
}

func insertBook(n, a, ph, py, id string, c1, c2, c3 uint16) {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Добавлено")

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `BOOK` (`name`, `auths`, `pub_house`, `pub_year`, `count_in_hall1`, `count_in_hall2`, `count_in_hall3`, `id`) VALUES('%s', '%s', '%s', '%s', %d, %d, %d, '%s')", n, a, ph, py, c1, c2, c3, id)) //'2005-08-25'
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}

func deleteBook(id string) {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Удалено")

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("DELETE FROM `BOOK` WHERE `BOOK`.`id` = '%s'", id))
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}

func deleteUser(id uint16) {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Удалено")

	defer db.Close()

	select_var, err := db.Query(fmt.Sprintf("SELECT `hall` FROM `USER` WHERE `USER`.`id` = %d", id))
	if err != nil {
		panic(err)
	}

	var current_hall int

	for select_var.Next() {
		err = select_var.Scan(&current_hall)
		if err != nil {
			panic(err)
		}

		fmt.Println(current_hall)
	}

	defer select_var.Close()

	select_current_var, err := db.Query(fmt.Sprintf("SELECT `capacity` FROM `HALL` WHERE `HALL`.`number` = %d", current_hall))
	if err != nil {
		panic(err)
	}
	var current_count int

	for select_current_var.Next() {
		err = select_current_var.Scan(&current_count)
		if err != nil {
			panic(err)
		}
	}

	defer select_current_var.Close()

	plus_var, err := db.Query(fmt.Sprintf("UPDATE `HALL` SET `capacity` = %d WHERE `HALL`.`number` = %d", current_count+1, current_hall)) //'2005-08-25'
	if err != nil {
		panic(err)
	}

	defer plus_var.Close()

	insert, err := db.Query(fmt.Sprintf("DELETE FROM `USER` WHERE `USER`.`id` = '%d'", id))
	if err != nil {
		panic(err)
	}

	defer insert.Close()
}

func main() {
	var choose_1 int

	fmt.Println("Что вы хотите сделать?")
	fmt.Println("1 - Добавить новую книгу")                                                               // Добавление в таблицу BOOK новых книг
	fmt.Println("2 - Добавить нового читателя")                                                           // Добавление в таблицу USER новых людей
	fmt.Println("3 - Удалить книгу по шифру")                                                             // Удаление из таблицы BOOK записей на основе book_id
	fmt.Println("4 - Удалить читателя по номеру билета")                                                  // Удаление из таблицы USER записей на освнове user_id
	fmt.Println("5 - Узнать какие книги взял читатель и дату взятия")                                     // Доставание из таблицы USER данных book_1, book_2, book_3, book_1_date, book_2_date, book_3_date в строке с определенным user_id
	fmt.Println("6 - Как называется книга? (по шифру)")                                                   // Доставание из таблицы BOOK данных name в строке с определенным book_id
	fmt.Println("7 - Какой у книги шифр? (по названию)")                                                  // Доставание из таблицы BOOK данных book_id в строке с определенным name
	fmt.Println("8 - список читателей под удаление/перерегистрацию (те, кто записался более года назад)") // Доставание данных из таблицы USER в которых registration_date больше на 365 дней от текущей даты
	fmt.Println("9 - Закрепить книгу за читателем")                                                       // Изменение book_1/book2/book3 (в зависимости от того, какой слот свободный) и book_1_date, book_2_date, book_3_date в строке с определенным user_id
	fmt.Println("10 - Зафиксировать возвращение книги")                                                   // Удаление book_1/book2/book3 (в зависимости от того, в каком из слотов определенный book_id) и book_1_date, book_2_date, book_3_date в строке с определенным user_id
	fmt.Println("11 - Статистические данные")                                                             // Еще одно меню

	fmt.Scanln(&choose_1)

	switch choose_1 {
	case 1:
		scanner := bufio.NewScanner(os.Stdin)
		var n, a, ph, py, id string
		var c1, c2, c3 uint16

		fmt.Println("Введите название книги")
		if scanner.Scan() {
			n = scanner.Text()
		}
		fmt.Println("Введите автора книги")
		if scanner.Scan() {
			a = scanner.Text()
		}
		fmt.Println("Введите издательство")
		if scanner.Scan() {
			ph = scanner.Text()
		}
		fmt.Println("Введите дату издания (год-месяц-день)")
		fmt.Scanln(&py)
		fmt.Println("Введите количество книг завезенных в 1 зал")
		fmt.Scanln(&c1)
		fmt.Println("Введите количество книг завезенных в 2 зал")
		fmt.Scanln(&c2)
		fmt.Println("Введите количество книг завезенных в 3 зал")
		fmt.Scanln(&c3)
		fmt.Println("Придумайте книге шифр из 6 символов (A-Z, 0-9)")
		fmt.Scanln(&id)
		insertBook(n, a, ph, py, id, c1, c2, c3)
	case 2:
		scanner := bufio.NewScanner(os.Stdin)
		var last_name, date_birth, adress, academic_degree, admission_library string
		var passport_number, phone_number, hall uint32

		fmt.Println("Введите фамилию")
		fmt.Scanln(&last_name)
		fmt.Println("Введите дату рождения (год-месяц-день)")
		fmt.Scanln(&date_birth)
		fmt.Println("Введите адрес")
		if scanner.Scan() {
			adress = scanner.Text()
		}
		fmt.Println("Введите ученую степень (или ее отсутсвие)")
		if scanner.Scan() {
			academic_degree = scanner.Text()
		}
		fmt.Println("Введите сегодняшнее число")
		fmt.Scanln(&admission_library)
		fmt.Println("Введите номер паспорта")
		fmt.Scanln(&passport_number)
		fmt.Println("Введите номер телефона")
		fmt.Scanln(&phone_number)
		fmt.Println("Введите номер зала, из которого читатель будет брать книги")
		fmt.Scanln(&hall)
		fmt.Println(passport_number, phone_number)
		insertUser(last_name, date_birth, adress, academic_degree, admission_library, passport_number, phone_number, hall)
	case 3:
		var book_id string
		fmt.Println("Введите шифр книги, которую желаете удалить")
		fmt.Scanln(&book_id)
		deleteBook(book_id)
	case 4:
		var user_id uint16
		fmt.Println("Введите номер билета читателя, которого желаете удалить")
		fmt.Scanln(&user_id)
		deleteUser(user_id)
	case 5:
		var user_id uint16
		fmt.Println("Введите номер билета читателя, книги которого желаете узнать")
		fmt.Scanln(&user_id)
		selectUersBooks(user_id)
	case 6:
		var book_id string
		fmt.Println("Введите шифр книги, название которой желаете узнать")
		fmt.Scanln(&book_id)
		selectBooksName(book_id)
	case 7:
		scanner := bufio.NewScanner(os.Stdin)
		var book_name string
		fmt.Println("Введите название книги, шифр которой желаете узнать")
		if scanner.Scan() {
			book_name = scanner.Text()
		}
		fmt.Println(book_name)
		selectBooksID(book_name)
	case 8:
		fmt.Println("Потом")
	case 9:
		var user_id uint16
		var book_id string
		fmt.Println("Введите номер билета пользователя, которому желаете приписать книгу")
		fmt.Scanln(&user_id)
		fmt.Println("Введите шифр книги котрую необходимо добавить")
		fmt.Scanln(&book_id)
		replaceBookInUsers(book_id, user_id)
	case 10:
		var user_id uint16
		var a uint16
		fmt.Println("Введите номер билета пользователя, который желает вернуть книгу")
		fmt.Scanln(&user_id)
		fmt.Println("Введите номер книги котрую необходимо вернуть")
		fmt.Scanln(&a)
		returnBook(user_id, a)
	case 11:
		fmt.Println("Потом")
	}
}
