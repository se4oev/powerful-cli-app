### Exercises:
* update the walk tool so that it allows the user to provide more than one file extension
* improve the walk tool by adding more filtering options, such as files modified after a certain date or files with long file names
* create a companion tool for walk that restores the archived files in case they are needed again. Recreate the original directory by using the same approach you used to create the destination directory in the archiveFile() function. Then use the gzip.Reader type from the compress/gzip package to uncompress the archive files