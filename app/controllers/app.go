package controllers

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
	"github.com/revel/revel/session"
)

type App struct {
	*revel.Controller
}

//Render index page (Login page)
func (c App) Index() revel.Result {
	return c.Render()
}

//login page (Post) 
//Check username password/ create cookie for login
func (c App) LoginPost(inputEmail string, inputPassword string) revel.Result {

	//Stuct to hold db values
	type customerDetails struct {
		custID 		string
		password 	string
		fName 		string
		lName 		string
		isActive 	int
		hasReserved int
	}

	var cusDetails customerDetails
	//open connection to db
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT cusID, password, firstName, lastName, isActive, hasReserved FROM customer WHERE email = ?", inputEmail)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&cusDetails.custID, &cusDetails.password, &cusDetails.fName, &cusDetails.lName, &cusDetails.isActive, &cusDetails.hasReserved)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	db.Close()
	if cusDetails.password == inputPassword{
		if cusDetails.isActive == 1 {
			//Create Cookie
			session.InitSession(revel.RevelLog)
			c.Session.ID()
			c.Session.Set("loginTest", true)
			c.Session.Set("id", cusDetails.custID)
			c.Session.Set("isActive", cusDetails.isActive)
			c.Session.Set("hasReserved", cusDetails.hasReserved)
			//Redirect to login page
			if cusDetails.hasReserved == 1 {
				return c.Redirect(App.ViewReserved)
			} else {
				return c.Redirect(App.Dashboard)
			}
				
		} else {
			//Show error message
			c.Flash.Error("Account is marked as not active, please contact them to activate it")
			c.FlashParams()
			//Redirect to login page
			return c.Redirect(App.Index)
		}
	}else {
		//Show error message
		c.Flash.Error("Login in failed. Wrong Username/Password")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

//function to check if user has already logged in
func (c App) LoginCheck() bool {
	//get elements in cookie, if not found || logintest != true || both, return false
	results, err := c.Session.Get("loginTest")
	if err != nil {
		return false
	}
	_, err2 := c.Session.Get("id")
	if err2 != nil {
		return false
	}
	_, err3 := c.Session.Get("isActive")
	if err3 != nil {
		return false
	}

	if results == true {
		return true
	} else {
		return false
	}
}

//Logout 
func (c App) Logout() revel.Result {
	if c.LoginCheck() == true {
		delete(c.Session, "id") 
		delete(c.Session, "loginTest")
		delete(c.Session, "isActive")
		delete(c.Session, c.Session.ID())

		for k := range c.Session {
			delete(c.Session, k)
		}
		fmt.Println(c.Session)
		return c.Redirect(App.Index)
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

func (c App) ForgotPassword() revel.Result {
	return c.Render()
}

func (c App) ForgotPasswordPost(inputEmail string) revel.Result {
	type resetDetails struct {
		ans 	string
		sQuest string
	}
	var resetPassword resetDetails
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT sQuestion, answer FROM customer WHERE email =?", inputEmail)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
	// for each row, scan the result into our tag composite object
		err = results.Scan(&resetPassword.sQuest, &resetPassword.ans)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	if resetPassword.sQuest != "" || resetPassword.ans != "" {
		sql := "UPDATE customer SET password = ?  WHERE email = ?"
		update, err2 := db.Exec(sql, resetPassword.ans, inputEmail)
		if err2 != nil {
			panic(err2.Error())
		} else {
			RowsA, err2 := update.RowsAffected()
			if err2 != nil {
				fmt.Println("RowsAffected Error", err)
			}
			//if update success
			if RowsA > 0 {
				fmt.Println("Account successfully updated")
			}
		}
		db.Close()
		c.Flash.Success("Password Reset is Successful. Password is changed to answer for Secret Question :  " + resetPassword.sQuest)
		c.Flash.Success(" Remember to change your password after logging in to secure your account")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	} else {
		c.Flash.Error("Username / password does not exist in database")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.ForgotPassword)
	}
}

//Display dashboard
func (c App) Dashboard() revel.Result {
	if c.LoginCheck() == true {
		return c.Render()
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

//Registration page (bonus)
func (c App) Registration() revel.Result {
	return c.Render()
}

//Registration page (post) register the user
func (c App) RegPost(inputEmail, Password, fname, lname, sQuestion, answer string) revel.Result{
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT cusID FROM customer WHERE email =?", inputEmail)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if results.Next() == true {
		c.Flash.Error("Username already exist")
		c.FlashParams()
		return c.Redirect(App.Registration)
	} else {
		sql := "INSERT INTO customer (email, password, firstName, lastName, isActive, hasReserved, sQuestion, answer) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
		insert, err2 := db.Query(sql, inputEmail, Password, fname, lname, "1" , 0, sQuestion, answer)
		if err2 != nil {
			panic(err2.Error())
		}
		insert.Close()
		db.Close()
		c.Flash.Success("Account successfully registered")
		c.FlashParams()
		return c.Redirect(App.Index)
	}
}

//Booking page
func (c App) Booking() revel.Result {
	if c.LoginCheck() == true {
		return c.Render()
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

//Display Bookinglist for booking form push out as (JSON)
func (c App) BookingList() revel.Result{
	if c.LoginCheck() == true {
		//Struct for flights
		type flights struct {
			ID 			string 	`json:"airlinesID"`
			Name 		string 	`json:"nameAL"`
			FlightNum	string	`json:"flightNumber"`
			FliesTo 	string 	`json:"flyingTo"`
			Date 		string 	`json:"flightDate"`
			Time 		string 	`json:"flightTime"`
			Seat 		string 	`json:"numberOfSeat"`
			Taken 		string 	`json:"seatTaken"`
			Reserved 	string 	`json:"reserved"`
		}
		var flightTemp flights
		var flightList []flights
	
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
		if err != nil {
			panic(err.Error())
		}
		results, err := db.Query("SELECT airlinesID,nameAL,flightNumber,flyingTo,flightDate,flightTime,numberOfSeat,seatTaken,reserved FROM airlines WHERE numberOfSeat != 0")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		for results.Next() {
		// for each row, scan the result into our tag composite object
			err = results.Scan(&flightTemp.ID, &flightTemp.Name, &flightTemp.FlightNum, &flightTemp.FliesTo, &flightTemp.Date, &flightTemp.Time, &flightTemp.Seat, &flightTemp.Taken, &flightTemp.Reserved)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			flightList = append(flightList, flights{ID: flightTemp.ID, Name: flightTemp.Name, FlightNum: flightTemp.FlightNum, FliesTo: flightTemp.FliesTo, Date: flightTemp.Date, 
				Time: flightTemp.Time, Seat: flightTemp.Seat, Taken :flightTemp.Taken, Reserved: flightTemp.Reserved})
			//fmt.Println(flightList)
		}
		if len(flightList) == 0 {
			flightTemp.ID = "-"
			flightTemp.Name = "-"
			flightTemp.FlightNum = "-"
			flightTemp.FliesTo = "-"
			flightTemp.Date = "-"
			flightTemp.Time = "-"
			flightTemp.Seat = "-"
			flightTemp.Taken = "-"
			flightTemp.Reserved = "-"
			flightList = append(flightList, flights{ID: flightTemp.ID, Name: flightTemp.Name, FlightNum: flightTemp.FlightNum, FliesTo: flightTemp.FliesTo, Date: flightTemp.Date, 
				Time: flightTemp.Time, Seat: flightTemp.Seat, Taken :flightTemp.Taken, Reserved: flightTemp.Reserved})
		}
		db.Close()
		fmt.Println(flightList)
		return c.RenderJSON(flightList)
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

//Render reserved booking page
func (c App) ViewReserved() revel.Result {
	if c.LoginCheck() == true {
		return c.Render()
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

//Display ReservedList for Reserved booking form push out as (JSON)
func (c App) BookingReserveList() revel.Result {
	if c.LoginCheck() == true {
		type reserved struct {
			ID 			string 	`json:"airlinesID"`
			Name 		string 	`json:"airlineName"`
			FlightNum	string	`json:"flightNumber"`
			FliesTo 	string 	`json:"flyingTo"`
			Date 		string 	`json:"bookingDate"`
			Time 		string 	`json:"bookingTime"`
			Seat 		string 	`json:"ticketAmount"`
		}
		//var flightTemp flights
		var reservedTemp reserved
		//var flightList []flights
		var reservedList []reserved

		cusID, err2 := c.Session.Get("id")
		if err2 != nil {
			panic(err2.Error())
		}

		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
		if err != nil {
			panic(err.Error())
		}
		results, err := db.Query("SELECT airlinesID,airlineName,flightNumber,flyingTo,bookingDate,bookingTime,ticketAmount FROM reserve WHERE cusID =?", cusID)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		for results.Next() {
		// for each row, scan the result into our tag composite object
			err = results.Scan(&reservedTemp.ID, &reservedTemp.Name, &reservedTemp.FlightNum, &reservedTemp.FliesTo, &reservedTemp.Date, &reservedTemp.Time, &reservedTemp.Seat)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			reservedList = append(reservedList, reserved{ID: reservedTemp.ID, Name: reservedTemp.Name, FlightNum: reservedTemp.FlightNum, 
				FliesTo: reservedTemp.FliesTo, Date: reservedTemp.Date, Time: reservedTemp.Time, Seat: reservedTemp.Seat})
			//fmt.Println(flightList)
		}
		if len(reservedList) == 0 {
			reservedTemp.ID = "-"
			reservedTemp.Name = "-"
			reservedTemp.FlightNum = "-"
			reservedTemp.FliesTo = "-"
			reservedTemp.Date = "-"
			reservedTemp.Time = "-"
			reservedTemp.Seat = "-"
			reservedList = append(reservedList, reserved{ID: reservedTemp.ID, Name: reservedTemp.Name, FlightNum: reservedTemp.FlightNum, 
				FliesTo: reservedTemp.FliesTo, Date: reservedTemp.Date, Time: reservedTemp.Time, Seat: reservedTemp.Seat})
			//return c.Redirect(App.Index) //need change to error page
		}
		db.Close()
		fmt.Println(reservedList)
		return c.RenderJSON(reservedList)
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

//Reserved booking form (Post) handle
func (c App) BookingReservePost(flightChoice, passportValues, nameValues, cardNum, formType string) revel.Result {
	fmt.Println(flightChoice)
	fmt.Println(passportValues)
	fmt.Println(nameValues)
	fmt.Println(formType)

	split := strings.Split(flightChoice, "@")
	seat, err := strconv.Atoi(split[len(split)-1])
	if err != nil{
		panic(err.Error())
	}
	var airlinesID string = split[0]
	var airlinesName string = split[1]
	var flightNumber string = split[2]
	var country string = split[3]
	var date string = split[4]
	var time string = split[5]

	//get customerID
	cusID, err2 := c.Session.Get("id")
	if err2 != nil {
		panic(err2.Error())
	}

	//Insert into database booking details
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
	if err != nil {
		panic(err.Error())
	}

	if formType == "Complete" {
		sql := "INSERT INTO booking (cusID, airlinesID, airlineName, flightNumber, flyingTo, bookingDate, bookingTime, ticketAmount, creditCardNum) VALUES (? ,?, ?, ?, ?, ?, ?, ?, ?)"
		insert, err2 := db.Query(sql, cusID, airlinesID, airlinesName,flightNumber, country, date, time, seat, cardNum)
		if err2 != nil {
			panic(err2.Error())
		}
		insert.Close()

		//Get total seat number left
		var seatNumber int 
		results, err := db.Query("SELECT numberOfSeat FROM airlines WHERE airlinesID = ?", airlinesID)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		for results.Next() {
		// for each row, scan the result into our tag composite object
			err = results.Scan(&seatNumber)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
		}
		seatNumber = seatNumber + 1
		splitPPvalues := strings.Split(passportValues, "@")
		splitNameValues := strings.Split(nameValues, "@")
		for i := 0; i < len(splitPPvalues); i++ {
			seatString := strconv.Itoa(seatNumber)
			var ticketNum string = flightNumber + "-" + seatString
			sql := "INSERT INTO tickets (cusID, airlinesID, airlineName, ticketNum, flightNumber, flyingTo, passportNum, flightDate, flightTime, seatNumber, custName) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
			insert, err2 := db.Query(sql, cusID, airlinesID, airlinesName, ticketNum, 
				flightNumber, country, splitPPvalues[i], date, time, seatNumber, splitNameValues[i])
			if err2 != nil {
				panic(err2.Error())
			}
			insert.Close()
			seatNumber++;
		}

		//update database
		sql = "UPDATE airlines SET seatTaken = seatTaken + ?, reserved = reserved - ? WHERE airlinesID = ?"
		update, err2 := db.Exec(sql, seat, seat, airlinesID)
		if err2 != nil {
			panic(err2.Error())
		} else {
			RowsA, err2 := update.RowsAffected()
			if err2 != nil {
				fmt.Println("RowsAffected Error", err)
			}
			//if update success
			if RowsA > 0 {
				fmt.Println("Account successfully updated")
			}
		}
	} else {
		//update database add the seats back from reserved to number of seat
		sql := "UPDATE airlines SET numberOfSeat = numberOfSeat + ?, reserved = reserved - ? WHERE airlinesID = ?"
		update, err2 := db.Exec(sql, seat, seat, airlinesID)
		if err2 != nil {
			panic(err2.Error())
		} else {
			RowsA, err2 := update.RowsAffected()
			if err2 != nil {
				fmt.Println("RowsAffected Error", err)
			}
			//if update success
			if RowsA > 0 {
				fmt.Println("Account successfully updated")
			}
		}
	}

	//Delete reserved row by customer since he/she has booked
	delete, err := db.Query("DELETE FROM reserve WHERE cusID = ? && airlinesID = ? && flightNumber = ? && bookingDate =?", cusID, airlinesID, flightNumber, date)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Account succesfully deleted")
	}

	//check if reserved table still has values from customer
	results, err := db.Query("SELECT idreserve FROM reserve WHERE cusID = ?", cusID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if results.Next() == true {
		delete.Close()
		db.Close()
		return c.Redirect(App.Dashboard)
	} else {
		sql := "UPDATE customer SET hasReserved = ? WHERE cusID = ?"
		//execute update Exec to edit account information
		update, err2 := db.Exec(sql, 0, cusID)
		if err2 != nil {
			panic(err2.Error())
		} else {
			RowsA, err2 := update.RowsAffected()
			if err2 != nil {
				fmt.Println("RowsAffected Error", err)
			}
			//if update success
			if RowsA > 0 {
				fmt.Println("Account successfully updated")
			}
		}
		delete.Close()
		db.Close()
		return c.Redirect(App.Dashboard)
	}
}

//Booking form (Post) handle depending on button clicked (Reserve or Book)
func (c App) BookingPost(flightChoice, seatNumber string) revel.Result {
	if c.LoginCheck() == true {
		fmt.Println(flightChoice)
		fmt.Println(seatNumber)
		split := strings.Split(flightChoice, "@")
		seatAva, err := strconv.Atoi(split[6])
		if err != nil{
			panic(err.Error())
		}
		seatNumber, err := strconv.Atoi(seatNumber)
		if err != nil{
			panic(err.Error())
		}
		if seatNumber <= seatAva {
			var airlinesID string = split[0]
			var airlinesName string = split[1]
			var flightNumber string = split[2]
			var country string = split[3]
			var date string = split[4]
			var time string = split[5]
			
			//get customerID
			cusID, err2 := c.Session.Get("id")
			if err2 != nil {
				panic(err2.Error())
			}

			//Insert into database booking details
			db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
			if err != nil {
				panic(err.Error())
			}
			
			sql := "INSERT INTO reserve (cusID, airlinesID, airlineName, flightNumber, flyingTo, bookingDate, bookingTime, ticketAmount) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
			insert, err2 := db.Query(sql, cusID, airlinesID, airlinesName, flightNumber, country, date, time, seatNumber)
			if err2 != nil {
				panic(err2.Error())
			}
			insert.Close()
			
			var newSeatNumber = seatAva - seatNumber
			//update ticket number
			sql = "UPDATE airlines SET numberOfSeat = ?, reserved = reserved + ? WHERE airlinesID = ?"
			//execute update Exec to edit account information
			update, err2 := db.Exec(sql, newSeatNumber, seatNumber, airlinesID)
			if err2 != nil {
				panic(err2.Error())
			} else {
				RowsA, err2 := update.RowsAffected()
				if err2 != nil {
					fmt.Println("RowsAffected Error", err)
				}
				//if update success
				if RowsA > 0 {
					fmt.Println("Account successfully updated")
				}
			}
			sql = "UPDATE customer SET hasReserved = ? WHERE cusID = ?"
			//execute update Exec to edit account information
			update, err2 = db.Exec(sql, 1, cusID)
			if err2 != nil {
				panic(err2.Error())
			} else {
				RowsA, err2 := update.RowsAffected()
				if err2 != nil {
					fmt.Println("RowsAffected Error", err)
				}
				//if update success
				if RowsA > 0 {
					fmt.Println("Account successfully updated")
				}
			}

			db.Close()
			return c.Redirect(App.ViewReserved)
		} else {
			c.Flash.Error("Seat number is bigger than Seat available")
			c.FlashParams()
			return c.Redirect(App.Booking)
		}
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

//Display booked table at dashboard
func (c App) BookedList() revel.Result {
	if c.LoginCheck() == true {
		type booked struct {
			ID 			string 	`json:"airlinesID"`
			Name 		string 	`json:"airlineName"`
			FlightNum	string	`json:"flightNumber"`
			FliesTo 	string 	`json:"flyingTo"`
			Date 		string 	`json:"bookingDate"`
			Time 		string 	`json:"bookingTime"`
			Seat 		string 	`json:"ticketAmount"`
			BoughtWith  string	`json:"creditCardNum"`
		}
		//var flightTemp flights
		var BookedTicTemp booked
		//var flightList []flights
		var BookedTicList []booked

		cusID, err2 := c.Session.Get("id")
		if err2 != nil {
			panic(err2.Error())
		}

		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
		if err != nil {
			panic(err.Error())
		}
		results, err := db.Query("SELECT airlinesID,airlineName,flightNumber,flyingTo,bookingDate,bookingTime,ticketAmount,creditCardNum FROM booking WHERE cusID =?", cusID)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		for results.Next() {
		// for each row, scan the result into our tag composite object
			err = results.Scan(&BookedTicTemp.ID, &BookedTicTemp.Name, &BookedTicTemp.FlightNum, &BookedTicTemp.FliesTo, &BookedTicTemp.Date, &BookedTicTemp.Time, &BookedTicTemp.Seat, &BookedTicTemp.BoughtWith)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			BookedTicList = append(BookedTicList, booked{ID: BookedTicTemp.ID, Name: BookedTicTemp.Name, 
				FlightNum: BookedTicTemp.FlightNum, FliesTo: BookedTicTemp.FliesTo, Date: BookedTicTemp.Date, Time: BookedTicTemp.Time, Seat: BookedTicTemp.Seat, BoughtWith: BookedTicTemp.BoughtWith})
		}
		if len(BookedTicList) == 0 {
			BookedTicTemp.ID = "-"
			BookedTicTemp.Name = "-"
			BookedTicTemp.FlightNum = "-"
			BookedTicTemp.FliesTo = "-"
			BookedTicTemp.Date = "-"
			BookedTicTemp.Time = "-"
			BookedTicTemp.Seat = "-"
			BookedTicTemp.BoughtWith = "-"
			BookedTicList = append(BookedTicList, booked{ID: BookedTicTemp.ID, Name: BookedTicTemp.Name, 
				FlightNum: BookedTicTemp.FlightNum, FliesTo: BookedTicTemp.FliesTo, Date: BookedTicTemp.Date, Time: BookedTicTemp.Time, Seat: BookedTicTemp.Seat, BoughtWith: BookedTicTemp.BoughtWith})
		}
		db.Close()
		fmt.Println(BookedTicList)
		return c.RenderJSON(BookedTicList)
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

func (c App) TicketList() revel.Result {
	if c.LoginCheck() == true {
		type ticket struct {
			CustName	string 	`json:"custName"`
			Name 		string 	`json:"airlineName"`
			TicketNum 	string 	`json:"ticketNum"`
			FlightNum	string	`json:"flightNumber"`
			FliesTo 	string 	`json:"flyingTo"`
			PassNum 	string 	`json:"passportNum"`
			Date 		string 	`json:"flightDate"`
			Time 		string 	`json:"flightTime"`
			SeatNumber 	string 	`json:"seatNumber"`
		}
		//var flightTemp flights
		var ticketTemp ticket
		//var flightList []flights
		var ticketList []ticket

		cusID, err2 := c.Session.Get("id")
		if err2 != nil {
			panic(err2.Error())
		}

		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
		if err != nil {
			panic(err.Error())
		}
		results, err := db.Query("SELECT custName,airlineName,ticketNum,flightNumber,flyingTo,passportNum,flightDate,flightTime,seatNumber FROM tickets WHERE cusID =?", cusID)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		for results.Next() {
		// for each row, scan the result into our tag composite object
			err = results.Scan(&ticketTemp.CustName, &ticketTemp.Name, &ticketTemp.TicketNum, &ticketTemp.FlightNum, &ticketTemp.FliesTo, &ticketTemp.PassNum, &ticketTemp.Date, &ticketTemp.Time, &ticketTemp.SeatNumber)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			ticketList = append(ticketList, ticket{CustName: ticketTemp.CustName, Name: ticketTemp.Name, TicketNum: ticketTemp.TicketNum, FlightNum: ticketTemp.FlightNum, 
				FliesTo: ticketTemp.FliesTo, PassNum: ticketTemp.PassNum, Date: ticketTemp.Date, Time: ticketTemp.Time, SeatNumber: ticketTemp.SeatNumber})
			//fmt.Println(flightList)
		}
		if len(ticketList) == 0 {
			ticketTemp.CustName = "-"
			ticketTemp.Name = "-"
			ticketTemp.TicketNum = "-"
			ticketTemp.FlightNum = "-"
			ticketTemp.FliesTo = "-"
			ticketTemp.PassNum = "-"
			ticketTemp.Date = "-"
			ticketTemp.Time = "-"
			ticketTemp.SeatNumber = "-"
			ticketList = append(ticketList, ticket{CustName: ticketTemp.CustName, Name: ticketTemp.Name, TicketNum: ticketTemp.TicketNum, FlightNum: ticketTemp.FlightNum, 
				FliesTo: ticketTemp.FliesTo, PassNum: ticketTemp.PassNum, Date: ticketTemp.Date, Time: ticketTemp.Time, SeatNumber: ticketTemp.SeatNumber})
			//return c.Redirect(App.Index) //need change to error page
		}
		db.Close()
		fmt.Println(ticketList)
		return c.RenderJSON(ticketList)
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
} 

func (c App) ChangePass() revel.Result {
	if c.LoginCheck() == true {
		return c.Render()
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

func (c App) ChangePassPost(cPassword string) revel.Result {
	if c.LoginCheck() == true {
		cusID, err2 := c.Session.Get("id")
		if err2 != nil {
			panic(err2.Error())
		}
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/airlines")
		if err != nil {
			panic(err.Error())
		}
		sql := "UPDATE customer SET password = ?  WHERE cusID = ?"
		update, err2 := db.Exec(sql, cPassword, cusID)
		if err2 != nil {
			panic(err2.Error())
		} else {
			RowsA, err2 := update.RowsAffected()
			if err2 != nil {
				fmt.Println("RowsAffected Error", err)
			}
			//if update success
			if RowsA > 0 {
				fmt.Println("Account successfully updated")
			}
		}
		c.Flash.Success("Password Change Successful")
		c.FlashParams()
		return c.Redirect(App.Dashboard)
	}else{
		//Show error message
		c.Flash.Error("Kindly login first")
		c.FlashParams()
		//Redirect to login page
		return c.Redirect(App.Index)
	}
}

