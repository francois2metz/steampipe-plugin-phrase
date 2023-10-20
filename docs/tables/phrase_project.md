# Table: phrase_project

A project in Phrase handle keys and translations.

## Examples

### List project

```sql
select
  name
from
  phrase_project;
```

### List least recently updated project

```sql
select
  name
from
  phrase_project
order by
  updated_at;
```

### List project from one account

```sql
select
  name
from
  phrase_project
where
  account_id = 'oneaccountid';
```

### Get project by id

```sql
select
  name,
  created_at
from
  phrase_project
where
  id = 'oneprojectid';
```
