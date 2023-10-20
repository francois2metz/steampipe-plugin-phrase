# Table: phrase_member

A member in phrase has access to an account.

You **must** specify `account_id` in a where clause in order to use this table.

## Examples

### List members

```sql
select
  email,
  username,
  role
from
  phrase_member
where
  account_id = 'oneaccountid';
```

### List members by their least recently activity

```
select
  email,
  username,
  role
from
  phrase_member
where
  account_id = 'oneaccountid'
order by
  last_activity_at;
```

### List admin member

```
select
  email,
  username
from
  phrase_member
where
  account_id = 'oneaccountid'
  and role='Admin';
```


### Get member by id

```sql
select
  email,
  username
from
  phrase_member
where
  account_id = 'oneaccountid'
  and id = 'onememberid';
```
