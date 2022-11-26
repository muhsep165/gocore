# gocore
golang core template

Folder name use ``snake_case``, all case is lower

File name use ``camelCase``, except in ``features`` folder. 

Structure ``features`` explanation :
  1. Every time you create a Module/Feature, please create a folder using snake_case and make sure to create the appropriate word/short description for the module.
  2. Each module/feature must contain  ``requirefile`` the ``requirefile`` is an 
    - entity file,
    - controller file, 
    - service file, 
    - interface repo file, and 
    - implementation repo file.
  3. naming the filenames in this featured folder, following the rule:
    ``folder_name`` + ``.`` + ``requirefile``

For common function, find in helper/common.go or you can added a new one, if it doesn't already exist.

Function Names in each file must follow the rules
  1. Use ``CamelCase`` if it's public function or use ``camelCase`` if it's a private function.
  2. create the appropriate word/short description for the function.

