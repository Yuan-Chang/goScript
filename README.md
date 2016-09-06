- Function: quickly switch between multiple paths
- Initial setting
 1. set global path in ~/.bash_profile(point to goScript folder) and do "source ~/.bash_profile"
 2. change the content of the file named "global" in goScript folder, point it to goScript folder but not include '/goScript'
 3. give read/write access "chmod -R 777 goScript" folder
 4. create a folder and named it as "path" in goScript folder

- Usage
  1. . setpath "path name"
  2. . go "path name"
