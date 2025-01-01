# Go-Expense-Tracker-CLI
This is a project from Roadmap.sh, here is the link: https://roadmap.sh/projects/expense-tracker

The project currently keeps track of the budget and expenses with the use of .json files. The files in the project are populated with test data. They can be deleted and will be initialized again through running the project. The "expenses.json" file will be created once the first addition is made, and the "budget.json" file will be created once the first budget is set.

# Make file commands for Linux

# Make Build
# Make Install
Build and Install will allow you to simply use expense-tracker without the file location for faster commands by copying it to your /bin folder

# Make Uninstall
Uninstall will remove it from your /bin folder and allow you to build/install new changes made to the project



# For the following commands, I will assume you are on linux and have installed this into your /bin folder

# Add command, the description and category belong in "" while the price amount just needs to be a valid number, positive value, without quotes
expense-tracker add --description "description" --amount ## --category "category"

# Delete command, allows the user to delete expenses by thier ID. The ID doesn't need to be in quotes, just a valid id number in your list
expense-tracker delete --id ##

# List command, allows you to list all or just some of your expenses based on their category
expense-tracker list
expense-tracker list --category "category"

# Summary command, allows you to see a summary of all charges, how much money you've spent in total, or just the ones that belong to a particular month. The month requires no quotes just a valid number(1-12).
expense-tracker summary
expense-tracker --month ##

# Budget commands, there are two commands which are related to budget. The first allows you to set a budget for a particular month, while the second allows you to simply check the budget for a provided month. No quotes needed, just valid numbers for thier fields
expense-tracker budget --month ## --amount ##

expense-tracker check-budget --month ##

# Export command, allows the user to export their expenses to a CSV file. The files name will be expenses.csv
expense-tracker export