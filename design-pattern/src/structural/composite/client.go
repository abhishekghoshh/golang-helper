package composite

func Main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1"}
	folder1.Add(file1)

	folder2 := &Folder{name: "Folder2"}
	folder2.Add(file2, file3, folder2)

	folder2.Search("rose")
}
