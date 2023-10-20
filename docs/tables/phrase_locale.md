# Table: phrase_locale

A locale defines a user language and contains language and country related parameters. Many languages can be added to project to reflect all the variants and country regions to supported by the software project.

You **must** specify `project_id` in a where clause in order to use this table.

## Examples

### List project locales

```sql
select
  name,
  code
from
  phrase_locale
where
  project_id = 'oneprojectid';
```

### Get default locale from a project

```
select
  name,
  code
from
  phrase_locale
where
  project_id = 'oneprojectid'
  and "default";
```

### List locale with unverified translations

```
select
  name,
  code
from
  phrase_locale
where
  project_id = 'oneprojectid'
  and statictics_translations_unverified_count > 0;
```

### Get locale by id

```sql
select
  name,
  code
from
  phrase_locale
where
  project_id = 'oneprojectid'
  and id = 'onelocaleid';
```
