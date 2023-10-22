# Table: phrase_tag

Tags can be attached to keys with meaningful labels to always keep them well organized. Generally, tags are useful to track which keys belong to a certain feature or section of a project to allow translating and reviewing more efficiently.

You **must** specify `project_id` in a where clause in order to use this table.

## Examples

### List project tags

```sql
select
  name,
  keys_count
from
  phrase_tag
where
  project_id = 'oneprojectid';
```

### Order tags by creation date

```
select
  name
from
  phrase_tag
where
  project_id = 'oneprojectid'
order by
  created_at;
```

### Order tags by least associated keys

```
select
  name
from
  phrase_tag
where
  project_id = 'oneprojectid'
order by
  keys_count;
```

### Get tag by name

```sql
select
  name,
  keys_count
from
  phrase_tag
where
  project_id = 'oneprojectid'
  and name = 'onetagname';
```
