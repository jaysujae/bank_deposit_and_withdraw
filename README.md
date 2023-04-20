# Sample project with temporal and encore

## Documents
- run temporalite https://github.com/temporalio/temporalite/blob/main/CONTRIBUTING.md
- run tigerbeetle https://github.com/tigerbeetledb/tigerbeetle#quickstart
- encore run

# Run locally
- make run-tigerbeetle-format  ( create DB file)
- make run-tigerbeetle-start ( run DB )
- make run-temporalite-start ( run temporal cluster )
- make run-tctl-register ( register temporal namespace )
- make run-create-accounts ( create accounts)

## Concepts
- Account 1 is Bank. Account 2 is customer.
- Everything is bank's POV.
- Starting with 1000 transfer from customer -> bank . ( customer deposited 1000 to bank. customer has 1000 credit, bank has 1000 deposits)
- Authorization : bank 'reserves' customer's credit
- Authorization expire : customer failed to give money. customer loses the reserved credit.
- Present : customer gave money to bank. Customer gets the reserved credit.

# Workflow
- When authorized, we run 'wait for present' workflow.
- if 100 seconds pass, it expires. ( selector with prsent & 100 seconds future)
- if present comes, then we try to find the transaction ID with the same amount of money, if there is we signal 'present' so that expire doesn't happen.