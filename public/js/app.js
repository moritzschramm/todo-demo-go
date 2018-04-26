var app = new Vue({
  el: '#app',
  data: {
    todos: [],
    text: "",
    maxId: 99,
    loading: true,
    orderSelection: "created_at"
  },
  created: function() {

  	var xhttp = new XMLHttpRequest();
	xhttp.open("POST", "/todos", true);
	xhttp.onreadystatechange = function(vm) {
		if (this.readyState == XMLHttpRequest.DONE && this.status == 200) {
			
			var todos = JSON.parse(this.responseText);
			todos.forEach(function(todo) {
				vm.todos.push({id: todo.id, note: todo.note, done: todo.done, created_at: todo.created_at, updated_at: todo.updated_at});
				if(todo.id >= vm.maxId) {
					vm.maxId = todo.id + 1;
				}
			});
			vm.loading = false;
		}
	}.bind(xhttp, this);
	xhttp.send();
  },
  mounted: function() {

	document.getElementById('note-input').focus();
  },
  watch: {
  	orderSelection: function(order) {
  		
  		this.sortNotes();
  		document.getElementById('note-input').focus();
  	}
  },
  methods: {
  	sortNotes: function() {

  		var order = this.orderSelection;
  		this.todos.sort(function(a, b) {
  			if(a[order] < b[order]) {
  				return -1;
  			} else if (a[order] > b[order]) {
  				return 1;
  			}
  			return 0;
  		});
  	},
  	addNote: function() {
  		if(this.text !== "") {

  			this.loading = true;

	  		var xhttp = new XMLHttpRequest();
	  		xhttp.open("POST", "/todo", true);
			xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
			xhttp.onreadystatechange = function(vm) {
			if (this.readyState == XMLHttpRequest.DONE && this.status == 200) {
					
					var todo = JSON.parse(this.responseText);
					vm.todos.push({id: todo.id, note: todo.note, done: todo.done, created_at: todo.created_at, updated_at: todo.updated_at});
			  		vm.text = "";
			  		vm.sortNotes();
					vm.to
					if(todo.id >= vm.maxId) {
						vm.maxId = todo.id + 1;
					}
					vm.loading = false;
				}
			}.bind(xhttp, this);
	  		xhttp.send(encodeTodo({note: this.text, done: false}));
  		}
  		document.getElementById('note-input').focus();
  	},
  	editNote: function(todo) {

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
	  	this.sortNotes();
  	},
  	deleteNote: function(todo, index) {
  		this.todos.splice(index, 1);

  		var xhttp = new XMLHttpRequest();
	  	xhttp.open("POST", "/delete/todo/"+todo.id, true);
	  	xhttp.send();
	  	document.getElementById('note-input').focus();
  	}
  }
});

function encodeTodo(todo) {

	return "note="+todo.note+"&done="+todo.done;
}