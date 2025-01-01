# Go-Expense-Tracker-CLI
This is a project from Roadmap.sh, here is the link: https://roadmap.sh/projects/expense-tracker

The project currently keeps track of the budget and expenses with the use of .json files. The files in the project are populated with test data. They can be deleted and will be initialized again through running the project. The "expenses.json" file will be created once the first addition is made, and the "budget.json" file will be created once the first budget is set.

# --Make file commands for Linux--

# Make Build
Builds the project 
# Make Install
Moves the project into ./bin folder so you can simply use expense-tracker in terminal instead of the projects file location

# Make Uninstall
Uninstall will remove it from your ./bin folder and allow you to build/install new changes made to the project

# --Project Commands--
I will assume you are on linux and have installed this project into your ./bin folder

# expense-tracker add --description "description" --amount ## --category "category"
Add command, the description and category belong in "" while the price amount just needs to be a valid number, positive value, without quotes

# expense-tracker delete --id ##
Delete command, allows the user to delete expenses by thier ID. The ID doesn't need to be in quotes, just a valid id number in your list

# expense-tracker list
# expense-tracker list --category "category"
List command, allows you to list all or just some of your expenses based on their category

# expense-tracker summary
# expense-tracker summary --month ##
Summary command, allows you to see a summary of all charges, how much money you've spent in total, or just the ones that belong to a particular month. The month requires no quotes just a valid number(1-12).

# expense-tracker budget --month ## --amount ##
# expense-tracker check-budget --month ##
Budget commands, there are two commands which are related to budget. The first allows you to set a budget for a particular month and amount, while the second allows you to simply check the budget for a provided month. No quotes needed, just valid numbers for thier fields

# expense-tracker export
Export command, allows the user to export their expenses to a CSV file. The files name will be expenses.csv