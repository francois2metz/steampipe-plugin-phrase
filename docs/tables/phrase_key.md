# Table: phrase_key

Keys are used to identify translatable text strings within software code.

You **must** specify `project_id` in a where clause in order to use this table.

## Examples

### List project keys

```sql
select
  id,
  name,
  description
from
  phrase_key
where
  project_id = 'oneprojectid';
```

### List keys from a branch (entreprise only)

```sql
select
  name,
  description,
from
  phrase_key
where
  project_id = 'oneprojectid'
  and branch = 'my-feature-branch';
```

### List keys that match a string

```sql
select
  name,
  description,
from
  phrase_key
where
  project_id = 'oneprojectid'
  and q = 'mykey*'
```

### Get key by id

```sql
select
  name,
  description
from
  phrase_key
where
  project_id = 'oneprojectid'
  and id = 'onekeyid';
```
