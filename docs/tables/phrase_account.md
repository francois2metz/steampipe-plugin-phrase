# Table: phrase_account

An account in phrase is an organization that handle projects, members, translations.

## Examples

### List account

```sql
select
  name,
  company
from
  phrase_account;
```

### Get one account

```sql
select
  name,
  company
from
  phrase_account
where
  id = 'oneaccountid';
```
