

var app = new Vue({
  el: '#app',
  data: {
    todos: [],
    text: ""
  },
  created: function() {

  	var xhttp = new XMLHttpRequest();
	xhttp.open("POST", "/todos", true);
	xhttp.onreadystatechange = function(vm) {
		if (this.readyState == XMLHttpRequest.DONE && this.status == 200) {
			
			var todos = JSON.parse(this.responseText);
			todos.forEach(function(todo) {
				vm.todos.push({id: todo.id, note: todo.note, done: todo.done, inEdit: false});
			});
		}
	}.bind(xhttp, this);
	xhttp.send();
  },
  methods: {
  	addNote: function() {
  		if(this.text !== "") {
	  		this.todos.push({id: -1, note: this.text, done: false, inEdit: false});
	  		this.text = "";

	  		var xhttp = new XMLHttpRequest();
	  		xhttp.open("POST", "/todo", true);
			xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			xhttp.onreadystatechange = function(vm) {
			if (this.readyState == XMLHttpRequest.DONE && this.status == 200) {
					
					var todo = JSON.parse(this.responseText);
					vm.todos[vm.todos.length-1].id = todo.id;
				}
			}.bind(xhttp, this);
	  		xhttp.send(encodeTodo(this.todos[this.todos.length-1]));
  		}
  	},
  	editNote: function(todo) {

  		todo.inEdit = false;

  		this.editTodo(todo);
  	},
  	toggleTodo: function(todo) {

  		todo.done = !todo.done;

  		this.editTodo(todo);
  	},
  	editTodo: function(todo) {

  		var xhttp = new XMLHttpRequest();
	  	xhttp.open("POST", "/edit/todo/"+todo.id, true);
	  	xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
	  	xhttp.send(encodeTodo(todo));
  	},
  	deleteNote: function(todo, index) {
  		this.todos.splice(index, 1);

  		var xhttp = new XMLHttpRequest();
	  	xhttp.open("POST", "/delete/todo/"+todo.id, true);
	  	xhttp.send();
  	}
  }
})

function encodeTodo(todo) {

	return "note="+todo.note+"&done="+todo.done;
}