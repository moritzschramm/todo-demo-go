<template>
  <div class="columns is-centered">
    <div class="column is-half" id="app">

      <hr>
      <div class="level">
          <div class="level-left">
              <div class="level-item">
                  <span class="title">Current Todos</span>
              </div>
          </div>
          <div class="level-right">
              <order-select
                  v-bind:loader="loading"
                  v-model="orderSelection"
                  v-on:order-changed="orderChanged">
              </order-select>
          </div>
      </div>

      <hr>

      <transition-group name="todo-list" tag="div">       
                 
      <todo v-for="(todo, index) in todos"
          v-bind:key="todo.id"
          v-bind:todo="todo"
          v-bind:index="index"
          v-on:toggle="toggleTodo"
          v-on:edit="editTodo"
          v-on:delete="deleteTodo">
      </todo>

      <create-todo-button
          v-bind:key="maxId"
          v-on:create-todo="addTodo">
      </create-todo-button>

      </transition-group>

    </div>
  </div>
</template>

<script>
import Todo from './Todo.vue'
import OrderSelect from './OrderSelect.vue'
import CreateTodoButton from './CreateTodoButton.vue'

export default {
  name: 'app',
  data () {
    return {
      todos: [],
      maxId: 99,
      loading: true,
      orderSelection: "created_at"
    }
  },
  created () {

    // fetch all todos from server and populate list

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
  mounted () {

    document.getElementById('note-input').focus();
  },
  methods: {
    orderChanged (order) {

        this.sortTodos();
        document.getElementById('note-input').focus();
    },
    sortTodos () {

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
    addTodo (text) {
        if(text !== "") {

            this.loading = true;

            var xhttp = new XMLHttpRequest();
            xhttp.open("POST", "/todo", true);
            xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
            xhttp.onreadystatechange = function(vm) {
            if (this.readyState == XMLHttpRequest.DONE && this.status == 200) {
                    
                    var todo = JSON.parse(this.responseText);
                    vm.todos.push({id: todo.id, note: todo.note, done: todo.done, created_at: todo.created_at, updated_at: todo.updated_at});
                    vm.sortTodos();
                    if(todo.id >= vm.maxId) {
                        vm.maxId = todo.id + 1;
                    }
                    vm.loading = false;
                }
            }.bind(xhttp, this);
            xhttp.send(this.encodeTodo({note: text, done: false}));
        }
        document.getElementById('note-input').focus();
    },
    toggleTodo (todo) {

        todo.done = !todo.done;

        this.editTodo(todo);
    },
    editTodo (todo) {

        var xhttp = new XMLHttpRequest();
        xhttp.open("POST", "/edit/todo/"+todo.id, true);
        xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhttp.send(this.encodeTodo(todo));
        this.sortTodos();                                   // if name changed, sort todos
    },
    deleteTodo (event) {

        var todo = event.todo;
        var index = event.index;
        this.todos.splice(index, 1);                        // delete todo from frontend

        var xhttp = new XMLHttpRequest();
        xhttp.open("POST", "/delete/todo/"+todo.id, true);  // delete todo from backend
        xhttp.send();
        document.getElementById('note-input').focus();
    }, 
    encodeTodo (todo) {

        return "note="+todo.note+"&done="+todo.done;
    }
  },
  components: {
    'todo': Todo,
    'order-select': OrderSelect,
    'create-todo-button': CreateTodoButton
  }
}
</script>

<style>
/* custom classes to overwrite some bulma settings */
.is-static {
  margin-left: 11px;
  background-color: transparent; 
}
.line-through {
  text-decoration: line-through;
}

.level:last-child {
  margin-bottom: 1.5rem;
}

/* classes for transition animations */
.todo-list-item {
  transition: all .2s;
}

.todo-list-enter {
  opacity: 1;
}
.todo-list-leave-to {
  opacity: 0;
  transform: translateX(1rem);
}
.todo-list-leave-active {
  position: absolute;
}
.todo-list-leave-active .button {
  opacity: 0;
}
.todo-list-move {
  transition: transform .15s;
}
</style>
