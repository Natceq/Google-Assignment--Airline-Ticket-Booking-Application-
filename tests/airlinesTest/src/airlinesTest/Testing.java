package airlinesTest;

import java.io.FileWriter;
import java.util.ArrayList;
import java.util.List;

import org.openqa.selenium.By;
import org.openqa.selenium.ElementNotInteractableException;
import org.openqa.selenium.NoSuchElementException;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.WebElement;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.support.ui.Select;

public class Testing 
{
	public static void main(String[] args) throws NoSuchElementException, InterruptedException
	{
		//Directory of chromedriver
		System.setProperty("webdriver.chrome.driver", "C:\\Users\\Nathaniel\\Desktop\\Testing\\chromedriver_win32\\chromedriver.exe");
		//create object
		WebDriver driver = new ChromeDriver();
		//go to webpage (local host)
		//driver.manage().window().maximize();
		
		//Common variables
		ArrayList <String> testCollection = new ArrayList<String>();
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		String empty = "";
		String result = "";
		
		//RegisterTest
		String RegisterTest = "";
		String newEmail = "1212@gmail.com";
		String newPw = "456456";
		String newConfirmPW = "456456";
		String fName = "nat";
		String lName = "chan";
		String wrongRegEmail = "hahaha.com";
		String wrongPw = "121212";
		String shortPW = "121";
		
		String sAnswer1 = "michelle";
		String sAnswer2 = "nathaniel";
		int index1 = 1;
		int index2 = 0;
		
		result = registerTest(driver, empty, empty, empty, empty, empty, index1, empty, negativeTestcase);
		RegisterTest = "Register Test 1 (Empty all fields)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, empty, newPw, newConfirmPW, fName, lName, index1, sAnswer1, negativeTestcase);
		RegisterTest = "Register Test 2 (Empty email)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, empty, newConfirmPW, fName, lName, index2, sAnswer2, negativeTestcase);
		RegisterTest = "Register Test 3 (Empty pasword)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, newPw, empty, fName, lName, index1, sAnswer1, negativeTestcase);
		RegisterTest = "Register Test 3 (Empty confirm pasword)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, newPw, newConfirmPW, empty, lName, index1, sAnswer1, negativeTestcase);
		RegisterTest = "Register Test 4 (Empty first name)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, newPw, newConfirmPW, fName, empty, index1, sAnswer1, negativeTestcase);
		RegisterTest = "Register Test 5 (Empty last name)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, newPw, newConfirmPW, fName, lName, index2, empty, negativeTestcase);
		RegisterTest = "Register Test 6 (Empty secret answer with first question)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, newPw, newConfirmPW, fName, lName, index1, empty, negativeTestcase);
		RegisterTest = "Register Test 7 (Empty secret answer with different question)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, shortPW, shortPW, fName, lName, index2, sAnswer2, negativeTestcase);
		RegisterTest = "Register Test 8 (short but matching password)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, newPw, wrongPw, fName, lName, index1, sAnswer1, negativeTestcase);
		RegisterTest = "Register Test 9 (not matching password)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, wrongRegEmail, newPw, newConfirmPW, fName, lName, index1, sAnswer1, negativeTestcase);
		RegisterTest = "Register Test 10 (Wrong email format)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, newPw, newConfirmPW, fName, lName, index1, sAnswer2, postiveTestcase);
		RegisterTest = "Register Test 11 (Positive Test Case)  " + result;
		testCollection.add(RegisterTest);
		
		result = registerTest(driver, newEmail, newPw, newConfirmPW, fName, lName, index1, sAnswer2, negativeTestcase);
		RegisterTest = "Register Test 12 (Email/username already exist in database)  " + result;
		testCollection.add(RegisterTest);
		//System.out.println(testCollection);
		
				
		
		//Login Test
		String correctEmail = newEmail;
		String correctPW = newPw;
		String wrongEmail = "haha@gmail.com";
		String wrongPW = "000000";
		String loginTest = "";
		
		result = loginTest(driver, correctEmail, wrongPW, negativeTestcase);
		loginTest = "Login Test 1 (Correct Username and Wrong password)  " + result;
		testCollection.add(loginTest);
		
		result = loginTest(driver, wrongEmail, correctPW, negativeTestcase);
		loginTest = "Login Test 2 (Wrong Username and Correct password)  " + result;
		testCollection.add(loginTest);
		
		result = loginTest(driver, wrongEmail, wrongPW, negativeTestcase);
		loginTest = "Login Test 3 (Wrong Username and Wrong password)  " + result;
		testCollection.add(loginTest);
		
		result = loginTest(driver, empty, correctPW, negativeTestcase);
		loginTest = "Login Test 4 (Empty Username and Correct password)  " + result;
		testCollection.add(loginTest);
		
		result = loginTest(driver, correctEmail, empty, negativeTestcase);
		loginTest = "Login Test 5 (Correct Username and Empty password)  " + result;
		testCollection.add(loginTest);
		
		result = loginTest(driver, correctEmail, empty, negativeTestcase);
		loginTest = "Login Test 6 (Empty Username and Empty password)  " + result;
		testCollection.add(loginTest);
		
		result = loginTest(driver, correctEmail, correctPW, postiveTestcase);
		loginTest = "Login Test 7 (Correct Username and Correct password)  " + result;
		testCollection.add(loginTest);

		//Logout test
		result = logoutTest(driver, postiveTestcase);
		String logoutTest = "Logout Test 1 (Correct Login credentials and logout)  " + result;
		testCollection.add(logoutTest);
				
		//Route Accessing Test : Accessing page without login
		//Display pages
		String dashboard = "http://localhost:9000/dashboard";
		String ticketBooking = "http://localhost:9000/booking";
		String confirmBooking = "http://localhost:9000/reserved";
		String changePassword = "http://localhost:9000/changepassword";
		
		//GET/POST Request
		String logout = "http://localhost:9000/logout";
		String getBookingList = "http://localhost:9000/booking/list";
		String bookingPost = "http://localhost:9000/booking";
		String getbookingReservedList = "http://localhost:9000/booking/reserve";
		String reservePost = "http://localhost:9000/reserved";
		String bookedList = "http://localhost:9000/booked";
		String ticketList = "http://localhost:9000/ticketlist";
		
		result = routeTest(driver, dashboard, negativeTestcase);
		String routeTest = "Route Test 1 (Access Dashboard page without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, ticketBooking, negativeTestcase);
		routeTest = "Route Test 2 (Access Booking page without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, confirmBooking, negativeTestcase);
		routeTest = "Route Test 3 (Access Transcation page without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, changePassword, negativeTestcase);
		routeTest = "Route Test 4 (change password page without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, logout, negativeTestcase);
		routeTest = "Route Test 5 (Access logout Request:(GET) without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, getBookingList, negativeTestcase);
		routeTest = "Route Test 6 (Access bookinglist Request:(GET) without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, bookingPost, negativeTestcase);
		routeTest = "Route Test 7 (Access booking Request:(POST) without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, getbookingReservedList, negativeTestcase);
		routeTest = "Route Test 8 (Access reserveList Request:(GET) without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, reservePost, negativeTestcase);
		routeTest = "Route Test 9 (Access reserve Request:(POST) without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, bookedList, negativeTestcase);
		routeTest = "Route Test 10 (Access bookedlist Request:(GET) without login)  " + result;
		testCollection.add(routeTest);
		
		result = routeTest(driver, ticketList, negativeTestcase);
		routeTest = "Route Test 11 (Access ticketlist Request:(GET) without login)  " + result;
		testCollection.add(routeTest);
		
		
		//Booking Test
		String bookingTest = "";
		String filter1 = "Singapore Airlines";
		String filter2 = "2021-07-10";
		String filter3 = "malaysia";
		int indexForBooking = 1;
		String maxTicketBooking = "5";
		String ticketBookingZero = "0";
		String ticketBookingInsane = "500";
		
		//all form fields empty
		result = bookingTest(driver, correctEmail, correctPW, empty, empty, empty, indexForBooking, empty, negativeTestcase);
		bookingTest = "Booking Test 1 (all form fields for booking empty)  " + result;
		testCollection.add(bookingTest);
		
		result = bookingTest(driver, correctEmail, correctPW, empty, empty, empty, indexForBooking, maxTicketBooking, negativeTestcase);
		bookingTest = "Booking Test 2 (Testing with no search filters)  " + result;
		testCollection.add(bookingTest);
		
		result = bookingTest(driver, correctEmail, correctPW, empty, filter2, filter3, indexForBooking, empty, negativeTestcase);
		bookingTest = "Booking Test 3 (Booking without entering ticket number)  " + result;
		testCollection.add(bookingTest);
		
		result = bookingTest(driver, correctEmail, correctPW, empty, filter2, filter3, indexForBooking, ticketBookingZero, negativeTestcase);
		bookingTest = "Booking Test 4 (ticket number input 0)  " + result;
		testCollection.add(bookingTest);
		
		result = bookingTest(driver, correctEmail, correctPW, empty, filter2, filter3, indexForBooking, ticketBookingInsane, negativeTestcase);
		bookingTest = "Booking Test 5 (ticket number input 500)  " + result;
		testCollection.add(bookingTest);
		
		result = bookingTest(driver, correctEmail, correctPW, empty, filter2, filter3, indexForBooking, maxTicketBooking, postiveTestcase);
		bookingTest = "Booking Test 6 (Correct Booking Test with the correct form fields)  " + result;
		testCollection.add(bookingTest);
		
		//Login into after booking. Not completing booking form, will let the user be pushed to the reserved page to complete
		String reservedLoginTest = "";
		result = loginReservedPage(driver, correctEmail, correctPW, postiveTestcase);
		reservedLoginTest = "Customer Has Reserved Login Test (Login after customer leaves after booking without payment)  " + result;
		testCollection.add(reservedLoginTest);
		
		//Reserve Test
		String resevedTest = "";
		String creditCardNum = "123456789012345";
		int indexOfSelect = 1;
		String passportNum = "A123456B";
		String name = "Customer";
		
		result = reservedTest(driver, correctEmail, correctPW, empty, 0, empty, empty, negativeTestcase);
		resevedTest = "Reseved Booking Test 1 (empty field for all)  " + result;
		testCollection.add(resevedTest);
		
		result = reservedTest(driver, correctEmail, correctPW, empty, indexOfSelect, passportNum, name, negativeTestcase);
		resevedTest = "Reseved Booking Test 2 (empty credit card number field)  " + result;
		testCollection.add(resevedTest);
		
		result = reservedTest(driver, correctEmail, correctPW, creditCardNum, 0, passportNum, name, negativeTestcase);
		resevedTest = "Reseved Booking Test 3 (Customer did not choose any ticket selection)  " + result;
		testCollection.add(resevedTest);
		
		result = reservedTest(driver, correctEmail, correctPW, creditCardNum, indexOfSelect, empty, name, negativeTestcase);
		resevedTest = "Reseved Booking Test 4 (passport number field empty)  " + result;
		testCollection.add(resevedTest);
		
		result = reservedTest(driver, correctEmail, correctPW, creditCardNum, indexOfSelect, passportNum, empty, negativeTestcase);
		resevedTest = "Reseved Booking Test 5 (name field empty)  " + result;
		testCollection.add(resevedTest);
		
		result = reservedTest(driver, correctEmail, correctPW, creditCardNum, indexOfSelect, passportNum, name, postiveTestcase);
		resevedTest = "Reseved Booking Test 6 (Correct flow with correct fields)  " + result;
		testCollection.add(resevedTest); 
		
		//Delete reserved test
		boolean click = true;
		boolean dontClick = false;
		String reservedDelete = "";
		
		//Create booking
		result = bookingTest(driver, correctEmail, correctPW, empty, filter2, filter3, indexForBooking, maxTicketBooking, postiveTestcase);
		
		result = reservedDeleteTest(driver, correctEmail, correctPW, 0, click, negativeTestcase);
		reservedDelete = "Reserved Delete Test 1 (Dont select user choice and submit for delete)  " + result;
		testCollection.add(reservedDelete);
		
		result = reservedDeleteTest(driver, correctEmail, correctPW, indexOfSelect, dontClick, negativeTestcase);
		reservedDelete = "Reserved Delete Test 2 (Dont check delete checkbox and submit for delete)  " + result;
		testCollection.add(reservedDelete);
		
		result = reservedDeleteTest(driver, correctEmail, correctPW, indexOfSelect, click, postiveTestcase);
		reservedDelete = "Reserved Delete Test 3 (Correct flow with user choice and checkbox click)  " + result;
		testCollection.add(reservedDelete);
		
		//Reset password test
		String resetPassword = "";
		
		result = resetPasswordTest(driver, empty, sAnswer2, negativeTestcase);
		resetPassword = "Reset Password Test 1 (empty email field)  " + result;
		testCollection.add(resetPassword);
		
		result = resetPasswordTest(driver, wrongEmail, sAnswer2, negativeTestcase);
		resetPassword = "Reset Password Test 2 (email address that does not exist in database)  " + result;
		testCollection.add(resetPassword);
		
		result = resetPasswordTest(driver, correctEmail, sAnswer2, postiveTestcase);
		resetPassword = "Reset Password Test 3 (Correct email and password)  " + result;
		testCollection.add(resetPassword);
		
		//Change password test
		String passwordChange  = "";
		
		result = changePasswordTest(driver, correctEmail, sAnswer2, empty, newConfirmPW, negativeTestcase);
		passwordChange = "Change Password Test 1 (empty password form field)  " + result;
		testCollection.add(passwordChange);
		
		result = changePasswordTest(driver, correctEmail, sAnswer2, newPw, empty, negativeTestcase);
		passwordChange = "Change Password Test 2 (empty confirm password form field)  " + result;
		testCollection.add(passwordChange);
		
		result = changePasswordTest(driver, correctEmail, sAnswer2, newPw, newConfirmPW, postiveTestcase);
		passwordChange = "Change Password Test 3 (Correct email and password)  " + result;
		testCollection.add(passwordChange);
		
		//End to end positive test case with a different account
		String newUserEmail = "newuser@gmail.com";
		String password = "hello123";
		filter2 = "2021-07-11";
		filter3 = "Hong kong";
		indexForBooking = 5;
		maxTicketBooking = "3";
		
		testCollection.add("End to end positive test case with a different account");
		result = registerTest(driver, newUserEmail, password, password, fName, lName, index1, sAnswer2, postiveTestcase);
		RegisterTest = "Register Test new different user (Positive Test Case)  " + result;
		testCollection.add(RegisterTest);
		
		result = loginTest(driver, newUserEmail, password, postiveTestcase);
		loginTest = "Login Test new different user (Correct Username and Correct password)  " + result;
		testCollection.add(loginTest);
		
		result = logoutTest(driver, postiveTestcase);
		logoutTest = "Logout Test new different user (Correct Login credentials and logout)  " + result;
		testCollection.add(logoutTest);
		
		result = bookingTest(driver, newUserEmail, password, empty, filter2, filter3, indexForBooking, maxTicketBooking, postiveTestcase);
		bookingTest = "Booking Test new different user (Correct Booking Test with the correct form fields)  " + result;
		testCollection.add(bookingTest);
		
		result = reservedTest(driver, newUserEmail, password, creditCardNum, indexOfSelect, passportNum, name, postiveTestcase);
		resevedTest = "Reseved Booking new different user (Correct flow with correct fields)  " + result;
		testCollection.add(resevedTest); 
		
		result = bookingTest(driver, newUserEmail, password, empty, filter2, filter3, indexForBooking, maxTicketBooking, postiveTestcase);
		result = reservedDeleteTest(driver, newUserEmail, password, indexOfSelect, click, postiveTestcase);
		reservedDelete = "Reserved Delete new different user (Correct flow with user choice and checkbox click)  " + result;
		testCollection.add(reservedDelete);
		
		result = resetPasswordTest(driver, newUserEmail, sAnswer2, postiveTestcase);
		resetPassword = "Reset Password new different user (Correct email and password)  " + result;
		testCollection.add(resetPassword);
		
		result = changePasswordTest(driver, newUserEmail, sAnswer2, password, password, postiveTestcase);
		passwordChange = "Change Password new different user (Correct email and password)  " + result;
		testCollection.add(passwordChange);
		
		
		driver.quit();
		//write to result file
		try 
		{
			FileWriter myWriter = new FileWriter("TestingResults.txt");
			for(String temp : testCollection)
			{
				myWriter.write(temp + "\n");
			}
			myWriter.close();
		}
		catch (Exception e)
		{
			e.printStackTrace();
		}
		//print out results
		for(String temp : testCollection)
		{
			System.out.println(temp);
		}
	}
	
	
	public static String logoutTest(WebDriver driver, String testCase) throws InterruptedException
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
		Thread.sleep(1500);
		driver.findElement(By.xpath("//a[@id = 'logout']")).click();
		Thread.sleep(1500);
		String correctTitle = "Login Page";
		String driverTitle = driver.getTitle();
		
		if(testCase.equals(postiveTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				results = "[" + postiveTestcase + "]   " + results + " PASSED";
			}
			else 
			{
				results = "[" + postiveTestcase + "]   " + results + " FAILED";
			}
		}
		if(testCase.equals(negativeTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				results = "[" + negativeTestcase + "]   " + results + " FAILED";
			}
			else 
			{
				results = "[" + negativeTestcase + "]   " + results + " PASSED";
			}
		}
		return results;
	}
	
	
	public static String routeTest(WebDriver driver, String url, String testCase) throws InterruptedException
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		driver.get(url);
		Thread.sleep(1500);
		String pushedOutTitle = "Login Page";
		String driverTitle = driver.getTitle();
		
		if(testCase.equals(negativeTestcase))
		{
			//user will be pushed out to login page when they try to access pages without logging in
			//if the login page == driver page, negative test case passed
			if(pushedOutTitle.equalsIgnoreCase(driverTitle) ==  false)
			{
				results = "[" + negativeTestcase + "]   " + results + " FAILED";
			}
			else 
			{
				results = "[" + negativeTestcase + "]   " + results + " PASSED";
			}
		}
		return results;
	}
	
	
	public static String registerTest(WebDriver driver, String email, String pw, String confirmPw, String firstName, String LastName, int index, String sAnswer,  String testCase) throws InterruptedException
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		
		driver.get("http://localhost:9000/registration");
		driver.findElement(By.id("inputEmail")).sendKeys(email);
		Thread.sleep(1000);
		driver.findElement(By.name("Password")).sendKeys(pw);
		Thread.sleep(1000);
		driver.findElement(By.id("cPassword")).sendKeys(confirmPw);
		Thread.sleep(1000);
		driver.findElement(By.id("fname")).sendKeys(firstName);
		Thread.sleep(1000);
		driver.findElement(By.id("lname")).sendKeys(LastName);
		Select sQuestion = new Select(driver.findElement(By.id("sQuestion")));
		sQuestion.selectByIndex(index);
		Thread.sleep(1000);
		driver.findElement(By.id("answer")).sendKeys(sAnswer);
		Thread.sleep(1000);
		driver.findElement(By.xpath("//button[@id = 'regSubmit']")).click();
		Thread.sleep(1500);
		String correctTitle = "Login Page";
		String driverTitle = driver.getTitle();
		//boolean exists1 = driver.findElements( By.id("displaySuccess") ).size() != 0;
		
		if(testCase.equals(postiveTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				results = "[" + postiveTestcase + "]   " + results + " PASSED";
			}
			else 
			{
				results = "[" + postiveTestcase + "]   " + results + " FAILED";
			}
		}
		if(testCase.equals(negativeTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				results = "[" + negativeTestcase + "]   " + results + " FAILED";
			}
			else 
			{
				results = "[" + negativeTestcase + "]   " + results + " PASSED";
			}
		}
		return results;
	}
	
	
	public static String loginTest(WebDriver driver, String email, String password, String testCase) throws InterruptedException 
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		driver.get("http://localhost:9000/");
		driver.findElement(By.id("inputEmail")).sendKeys(email);
		driver.findElement(By.name("inputPassword")).sendKeys(password);
		Thread.sleep(1500);
		driver.findElement(By.xpath("//button[@id = 'loginSubmit']")).click();
		Thread.sleep(1500);
		String correctTitle = "Customer Dashboard";
		String driverTitle = driver.getTitle();
		if(testCase.equals(postiveTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				results = "[" + postiveTestcase + "]   " + results + " PASSED";
			}
			else 
			{
				results = "[" + postiveTestcase + "]   " + results + " FAILED";
			}
		}
		if(testCase.equals(negativeTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				results = "[" + negativeTestcase + "]   " + results + " FAILED";
			}
			else 
			{
				results = "[" + negativeTestcase + "]   " + results + " PASSED";
			}
		}
		return results;
	}
	
	
	public static String bookingTest(WebDriver driver, String email, String password, String search1, String search2, String search3, int indexForTicket, String seatNumber, String testCase) throws ElementNotInteractableException,InterruptedException
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		try 
		{
			//Homepage login
			driver.get("http://localhost:9000/");
			driver.findElement(By.id("inputEmail")).sendKeys(email);
			driver.findElement(By.name("inputPassword")).sendKeys(password);
			Thread.sleep(1500);
			driver.findElement(By.xpath("//button[@id = 'loginSubmit']")).click();
			Thread.sleep(1500);
			//Enter booking page from dashboard
			Thread.sleep(1500);
			//Absolute path
			driver.findElement(By.xpath("/html/body/div[1]/div/div/div[2]/p[1]/a")).click();
			driver.findElement(By.id("myInput")).sendKeys(search1);
			driver.findElement(By.id("myInputDate")).sendKeys(search2);
			driver.findElement(By.id("myInputCountry")).sendKeys(search3);
			Thread.sleep(1500);
			Select sQuestion = new Select(driver.findElement(By.id("flightChoice")));
			sQuestion.selectByIndex(indexForTicket);
			Thread.sleep(1500);
			driver.findElement(By.id("seatNumber")).sendKeys(seatNumber);
			
			//Throw ElementNotInteractableException as it is set as hidden unless form values are entered.
			driver.findElement(By.xpath("/html/body/div[1]/div/div/div/div[2]/button")).click();
			Thread.sleep(1500);
			String driverTitle = driver.getTitle();
			String correctTitle = "Reserved Bookings";
			//Check Table row, if table row is not - means booking is successful
			WebElement rwData = driver.findElement(By.xpath("/html/body/div[1]/div/div/div/div/table/tbody[2]/tr/td[1]"));
			
			if(testCase.equals(postiveTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					//If booking record is saved, value in table will not be "-" if not "-" testcase passed
					if(rwData.getText().equals("-") == false)
					{
						results = "[" + postiveTestcase + "]   " + results + " PASSED";
					}
					else 
					{
						results = "[" + postiveTestcase + "]   " + results + " FAILED";
					}
				}
				else 
				{
					results = "[" + postiveTestcase + "]   " + results + " FAILED";
				}
			}
			if(testCase.equals(negativeTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + negativeTestcase + "]   " + results + " FAILED";
				}
				else 
				{
					results = "[" + negativeTestcase + "]   " + results + " PASSED";
				}
			}
			driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
			Thread.sleep(1500);
			driver.findElement(By.xpath("//a[@id = 'logout']")).click();
			
			return results;
		}
		catch (Exception e)
		{
			String driverTitle = driver.getTitle();
			String correctTitle = "Reserved Bookings";
			if(testCase.equals(postiveTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + postiveTestcase + "]   " + results + " PASSED";
				}
				else 
				{
					results = "[" + postiveTestcase + "]   " + results + " FAILED";
				}
			}
			if(testCase.equals(negativeTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + negativeTestcase + "]   " + results + " FAILED";
				}
				else 
				{
					results = "[" + negativeTestcase + "]   " + results + " PASSED";
				}
			}
			driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
			Thread.sleep(1500);
			driver.findElement(By.xpath("//a[@id = 'logout']")).click();
			return results;
		}
	}
	
	
	public static String loginReservedPage(WebDriver driver, String email, String password, String testCase) throws InterruptedException
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		//Homepage login
		driver.get("http://localhost:9000/");
		driver.findElement(By.id("inputEmail")).sendKeys(email);
		driver.findElement(By.name("inputPassword")).sendKeys(password);
		Thread.sleep(1500);
		driver.findElement(By.xpath("//button[@id = 'loginSubmit']")).click();
		Thread.sleep(1500);
		//Once user is logged in, user should be pushed into reserved bookings page.
		String driverTitle = driver.getTitle();
		String correctTitle = "Reserved Bookings";
		
		//Get table row value, if table row value != "-", table customer has reserved tickets and did not complete the transaction
		
		WebElement rwData = driver.findElement(By.xpath("/html/body/div[1]/div/div/div/div/table/tbody[2]/tr/td[1]"));
		if(testCase.equals(postiveTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				//If booking record is saved, value in table will not be "-" if not "-" testcase passed
				if(rwData.getText().equals("-") == false)
				{
					results = "[" + postiveTestcase + "]   " + results + " PASSED";
				}
				else 
				{
					results = "[" + postiveTestcase + "]   " + results + " FAILED";
				}
			}
			else 
			{
				results = "[" + postiveTestcase + "]   " + results + " FAILED";
			}
		}
		if(testCase.equals(negativeTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				results = "[" + negativeTestcase + "]   " + results + " FAILED";
			}
			else 
			{
				results = "[" + negativeTestcase + "]   " + results + " PASSED";
			}
		}
		driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
		Thread.sleep(1500);
		driver.findElement(By.xpath("//a[@id = 'logout']")).click();
		return results;
	}
	
	
	public static String reservedTest(WebDriver driver, String email, String password, String creditCardNum, int index, String ppNum, String name, String testCase) throws NoSuchElementException, InterruptedException
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		//driver.findElement(By.xpath("//*[@id=\"forDelete\"]")).click();
		driver.get("http://localhost:9000/");
		driver.findElement(By.id("inputEmail")).sendKeys(email);
		driver.findElement(By.name("inputPassword")).sendKeys(password);
		driver.findElement(By.xpath("//button[@id = 'loginSubmit']")).click();
		Thread.sleep(1000);
		Select userChoice = new Select(driver.findElement(By.id("flightChoice")));
		try
		{
			//Select choice
			userChoice.selectByIndex(index);
			
			//Get selected choice
			WebElement option = userChoice.getFirstSelectedOption();
			String selectedValue = option.getText();
			Thread.sleep(1000);
			//Get table row, search for number of seats
			String seatNumString = "";
			WebElement tableBody = driver.findElement(By.xpath("/html/body/div[1]/div/div/div/div/table/tbody[2]"));
		    List<WebElement> rows = tableBody.findElements(By.tagName("tr"));
		    for (WebElement row : rows) {
		        List<WebElement> td = row.findElements(By.tagName("td"));
		        if (td.size() > 0 && td.get(0).getText().equals(selectedValue)) 
		        {
		        	seatNumString = td.get(6).getText();
		        }
		    }
		    Thread.sleep(1000);
		    //get seat number for forLoop
		    int seatNum = Integer. parseInt(seatNumString);
			for(int i = 0; i < seatNum; i++)
			{
				Thread.sleep(1000);
				if(ppNum.equals(""))
				{
					driver.findElement(By.id("passNum" + i)).sendKeys(ppNum);
				}
				else
				{
					driver.findElement(By.id("passNum" + i)).sendKeys(ppNum + i);
				}
				if(name.equals(""))
				{
					driver.findElement(By.id("name" + i)).sendKeys(name);
				}
				else
				{
					driver.findElement(By.name("name" + i)).sendKeys(name + i);
				}
				Thread.sleep(1000);
			}
			Thread.sleep(1000);
			driver.findElement(By.xpath("//*[@id=\"cardNum\"]")).sendKeys(creditCardNum);
			driver.findElement(By.xpath("//*[@id=\"reservedSubmit\"]")).click();
			Thread.sleep(1500);
			String driverTitle = driver.getTitle();
			String correctTitle = "Customer Dashboard";
			//Check Table row, if table row is not - means transcation is successful
			WebElement rwData = driver.findElement(By.xpath("/html/body/div[1]/div/div/div[1]/table[1]/tbody[2]/tr/td[1]"));
			
			if(testCase.equals(postiveTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					//If booking record is saved, value in table will not be "-" if not "-" testcase passed
					if(rwData.getText().equals("-") == false)
					{
						results = "[" + postiveTestcase + "]   " + results + " PASSED";
					}
					else 
					{
						results = "[" + postiveTestcase + "]   " + results + " FAILED";
					}
				}
				else 
				{
					results = "[" + postiveTestcase + "]   " + results + " FAILED";
				}
			}
			if(testCase.equals(negativeTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + negativeTestcase + "]   " + results + " FAILED";
				}
				else 
				{
					results = "[" + negativeTestcase + "]   " + results + " PASSED";
				}
			}
			driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
			Thread.sleep(1500);
			driver.findElement(By.xpath("//a[@id = 'logout']")).click();
			return results;
		}
		catch (Exception e)
		{
			String driverTitle = driver.getTitle();
			String correctTitle = "Customer Dashboard";
			if(testCase.equals(postiveTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + postiveTestcase + "]   " + results + " PASSED";
				}
				else 
				{
					results = "[" + postiveTestcase + "]   " + results + " FAILED";
				}
			}
			if(testCase.equals(negativeTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + negativeTestcase + "]   " + results + " FAILED";
				}
				else 
				{
					results = "[" + negativeTestcase + "]   " + results + " PASSED";
				}
			}
			driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
			Thread.sleep(1500);
			driver.findElement(By.xpath("//a[@id = 'logout']")).click();
			return results;
		}
	}
	
	
	public static String reservedDeleteTest(WebDriver driver, String email, String password, int index, boolean click, String testCase) throws NoSuchElementException, InterruptedException
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		driver.get("http://localhost:9000/");
		driver.findElement(By.id("inputEmail")).sendKeys(email);
		driver.findElement(By.name("inputPassword")).sendKeys(password);
		driver.findElement(By.xpath("//button[@id = 'loginSubmit']")).click();
		Thread.sleep(1000);
		Select userChoice = new Select(driver.findElement(By.id("flightChoice")));
		try
		{
			//Select Choice
			if(index > 0)
				userChoice.selectByIndex(index);
			Thread.sleep(1000);
			if(click == true)
				driver.findElement(By.xpath("//*[@id=\"forDelete\"]")).click();
			Thread.sleep(1000);
			driver.findElement(By.xpath("//*[@id=\"reservedDelete\"]")).click();
			Thread.sleep(1500);
			
			//This throws NoSuchElementException as form did not go through, this element will be at the dashboard after form submit
			Thread.sleep(1000);
			driver.findElement(By.xpath("/html/body/div[1]/div/div/div[2]/p[2]/a")).click();
			Thread.sleep(1000);
			
			String driverTitle = driver.getTitle();
			String correctTitle = "Reserved Bookings";
			//Check Table row, if table row is not - means transcation is successful
			//WebElement rwData = driver.findElement(By.xpath("/html/body/div[1]/div/div/div[1]/table[1]/tbody[2]/tr/td[1]"));
			WebElement rwData = driver.findElement(By.xpath("/html/body/div[1]/div/div/div/div/table/tbody[2]/tr/td[1]"));
			if(testCase.equals(postiveTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					//If booking record is saved, value in table will not be "-" if not "-" testcase passed
					if(rwData.getText().equals("-"))
					{
						results = "[" + postiveTestcase + "]   " + results + " PASSED";
					}
					else 
					{
						results = "[" + postiveTestcase + "]   " + results + " FAILED";
					}
				}
				else 
				{
					results = "[" + postiveTestcase + "]   " + results + " FAILED";
				}
			}
			if(testCase.equals(negativeTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + negativeTestcase + "]   " + results + " FAILED";
				}
				else 
				{
					results = "[" + negativeTestcase + "]   " + results + " PASSED";
				}
			}
			driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
			Thread.sleep(1500);
			driver.findElement(By.xpath("//a[@id = 'logout']")).click();
			return results;
		}
		catch (Exception e)
		{
			String driverTitle = driver.getTitle();
			String correctTitle = "Reserved Bookings";
			WebElement rwData = driver.findElement(By.xpath("/html/body/div[1]/div/div/div/div/table/tbody[2]/tr/td[1]"));
			if(testCase.equals(postiveTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + postiveTestcase + "]   " + results + " PASSED";
				}
				else 
				{
					results = "[" + postiveTestcase + "]   " + results + " FAILED";
				}
			}
			if(testCase.equals(negativeTestcase))
			{
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					//If delete did not go through, value will not be table row value will not be -
					if(rwData.getText().equals("-") == false)
						results = "[" + negativeTestcase + "]   " + results + " PASSED";
					else
						results = "[" + negativeTestcase + "]   " + results + " FAILED";
				}
				else 
				{
					results = "[" + negativeTestcase + "]   " + results + " FAILED";
				}
			}
			driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
			Thread.sleep(1500);
			driver.findElement(By.xpath("//a[@id = 'logout']")).click();
			return results;
		}
	}

	
	public static String resetPasswordTest(WebDriver driver, String email, String answer, String testCase) throws InterruptedException
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		
		//go to reset password page
		driver.get("http://localhost:9000/forgotpassword");
		driver.findElement(By.id("inputEmail")).sendKeys(email);
		Thread.sleep(1500);
		driver.findElement(By.xpath("//button[@id = 'forgotPass']")).click();
		
		
		String correctTitle = "Login Page";
		String driverTitle = driver.getTitle();
		if(testCase.equals(postiveTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				//Login after reset
				driver.findElement(By.id("inputEmail")).sendKeys(email);
				driver.findElement(By.name("inputPassword")).sendKeys(answer);
				Thread.sleep(1500);
				driver.findElement(By.xpath("//button[@id = 'loginSubmit']")).click();
				Thread.sleep(1500);
				correctTitle = "Customer Dashboard";
				driverTitle = driver.getTitle();
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + postiveTestcase + "]   " + results + " PASSED";
				}
				else 
				{
					results = "[" + postiveTestcase + "]   " + results + " FAILED";
				}
				driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
				Thread.sleep(1500);
				driver.findElement(By.xpath("//a[@id = 'logout']")).click();
			}
			else 
			{
				results = "[" + postiveTestcase + "]   " + results + " FAILED";
			}
			
		}
		if(testCase.equals(negativeTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				results = "[" + negativeTestcase + "]   " + results + " FAILED";
			}
			else 
			{
				results = "[" + negativeTestcase + "]   " + results + " PASSED";
			}
		}
		return results;
	}


	public static String changePasswordTest(WebDriver driver, String email, String answer, String pw, String confirmPw, String testCase) throws InterruptedException
	{
		String results = "Results : ";
		String negativeTestcase = "NEGATIVE TESTCASE";
		String postiveTestcase = "POSTIVE TESTCASE";
		
		//Login
		driver.get("http://localhost:9000/");
		Thread.sleep(500);
		driver.findElement(By.id("inputEmail")).sendKeys(email);
		driver.findElement(By.name("inputPassword")).sendKeys(answer);
		Thread.sleep(1500);
		driver.findElement(By.xpath("//button[@id = 'loginSubmit']")).click();
		Thread.sleep(1500);
		//transverse to change password page
		driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
		Thread.sleep(1500);
		driver.findElement(By.xpath("//a[@id = 'changePasswordLink']")).click();
		
		//Pass values into form fields
		driver.findElement(By.name("Password")).sendKeys(pw);
		Thread.sleep(1000);
		driver.findElement(By.id("cPassword")).sendKeys(confirmPw);
		Thread.sleep(1000);
		driver.findElement(By.xpath("//button[@id = 'regSubmit']")).click();
		Thread.sleep(1500);
		String correctTitle = "Customer Dashboard";
		String driverTitle = driver.getTitle();
		if(testCase.equals(postiveTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				//Logout
				driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
				Thread.sleep(1500);
				driver.findElement(By.xpath("//a[@id = 'logout']")).click();
				
				//Login again
				driver.findElement(By.id("inputEmail")).sendKeys(email);
				driver.findElement(By.name("inputPassword")).sendKeys(pw);
				Thread.sleep(1500);
				driver.findElement(By.xpath("//button[@id = 'loginSubmit']")).click();
				Thread.sleep(1500);
				correctTitle = "Customer Dashboard";
				driverTitle = driver.getTitle();
				if(correctTitle.equalsIgnoreCase(driverTitle))
				{
					results = "[" + postiveTestcase + "]   " + results + " PASSED";
				}
				else 
				{
					results = "[" + postiveTestcase + "]   " + results + " FAILED";
				}
			}
			else 
			{
				results = "[" + postiveTestcase + "]   " + results + " FAILED";
			}
		}
		if(testCase.equals(negativeTestcase))
		{
			if(correctTitle.equalsIgnoreCase(driverTitle))
			{
				results = "[" + negativeTestcase + "]   " + results + " FAILED";
			}
			else 
			{
				results = "[" + negativeTestcase + "]   " + results + " PASSED";
			}
		}
		driver.findElement(By.xpath("//button[@id = 'custMenu']")).click();
		Thread.sleep(1500);
		driver.findElement(By.xpath("//a[@id = 'logout']")).click();
		return results;
	}
}
