function showHide(id){
	// show hide block element
	// REQUIRES that <el>.style.display be initialized to
	// "none" or "block" to work properly

	it = document.querySelector(id);
	if (it.style.display === "flex"){
		it.style.display = "none";
	} else {
		it.style.display = "flex";
	}
}