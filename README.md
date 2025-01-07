# Go-Expense-Tracker-CLI
![Alt text](images/Expense-Tracker.png)
This is a project from Roadmap.sh, here is the link: https://roadmap.sh/projects/expense-tracker

This project allows the user to keep track of expenses through the command line. You can add, delete, list, and summarize your expenses. You can also set a budget for a particular month. The budget is automatically checked when a new expense is added, but you can also manually  check with the check-budget command whenever you want. I have also implemented an export command to export the expenses to a CSV file.

The project currently keeps track of the budget and expenses with the use of .json files. The files in the project are populated with test data. They can be deleted and will be initialized again through running the project. The "expenses.json" file will be created once the first addition is made, and the "budget.json" file will be created once the first budget is set. The "expenses.csv" file will also be automatically generated by  the export command if it is deleted.

# Makefile commands for Linux

## Build
Builds the project 
```bash
make build
```
## Install
Moves the project into ./bin folder so you can simply use expense-tracker in terminal instead of the projects file location
```bash
make install
```

## Uninstall
Uninstall will remove it from your ./bin folder and allow you to build/install new changes made to the project
```bash
make uninstall
```

# Project Commands
 I will assume you are on linux and have installed this project into your ./bin folder

## Add
Adds an expense, placing the description and category belong in quotes while the price amount just needs to be a valid number, positive value, without quotes
```bash
expense-tracker add --description "Lunch" --amount 25 --category "Food"
```

## Delete
Allows the user to delete expenses by thier ID. The ID doesn't need to be in quotes, just a valid id number in your list
```bash
expense-tracker delete --id 2
```

## List 
Allows you to list all of their expenses
```bash
expense-tracker list
```
Allows the user to also filter the results by category, the category belongs in quotes
```bash
expense-tracker list --category "Food"
```

## Summary
Allows you to see a summary of all charges, the total amount they've spent
```bash
expense-tracker summary
```
Allows the user to filter the results by month, the month requires no quotes and must be a valid number (1-12)
```bash
expense-tracker summary --month 12
```

## Budget 
Allows you to set a budget for a particular month and amount, the month should be a valid number (1-12) and the amount should be a positive number
```bash
expense-tracker budget --month 12 --amount 200
```

## Check-Budget
Allows you to simply check the budget for a provided month, the month should be a valid number (1-12)
```bash
expense-tracker check-budget --month 12
```

## Export
Allows the user to export their expenses to a CSV file. The files name will be expenses.csv
```bash
expense-tracker export
```