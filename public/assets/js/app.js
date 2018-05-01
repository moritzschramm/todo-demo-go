Vue.component('create-todo-button', {
    data: function() {
        return {
            text: "",
        };
    },
    methods: {
        emitCreateTodo: function() {
            if(this.text != "") {

                this.$emit('create-todo', this.text);
                this.text = "";
            }
        }
    },
    template: '<div class="todo-list-item level is-mobile">                                                         \
                <div class="level-left">                                                                            \
                    <div class="level-item">                                                                        \
                        <label class="checkbox"><input type="checkbox" disabled></label>                            \
                    </div>                                                                                          \
                    <div class="leve-item">                                                                         \
                        <input id="note-input" class="input is-medium is-static" type="text" placeholder="New Todo" \
                            v-model.trim="text"                                                                     \
                            v-on:keydown.enter="emitCreateTodo">                                                    \
                    </div>                                                                                          \
                </div>                                                                                              \
                <div class="level-right">                                                                           \
                    <div class="level-item">                                                                        \
                        <button class="button is-success is-medium" v-on:click="emitCreateTodo">                    \
                            <span class="icon">                                                                     \
                                <i class="mdi mdi-24px mdi-plus"></i>                                               \
                            </span>                                                                                 \
                            <span>Create</span>                                                                     \
                        </button>                                                                                   \
                    </div>                                                                                          \
                </div>                                                                                              \
            </div>                                                                                                  \
    '
});

Vue.component('order-select', {
    props: {
        loader: Boolean,
        value: {
            type: String,
            default: "created_at"
        }
    },
    methods: {
        selectChanged: function(order) {
            this.$emit('input', order);
            this.$emit('order-changed', order);
        }
    },
    template: '<div class="level-item">                                                             \
                    <span style="padding-right: 10px;">Order by: </span>                            \
                    <div class="select"                                                             \
                            v-bind:class="{\'is-loading\': loader}">                                \
                        <select v-model="value" v-on:input="selectChanged($event.target.value)">    \
                            <option value="created_at">Date</option>                                \
                            <option value="note">Name</option>                                      \
                        </select>                                                                   \
                    </div>                                                                          \
                </div>'
});

Vue.component('todo', {
    props: ["todo", "index"],
    methods: {
        emitToggle: function() {

            this.$emit('toggle', this.todo);
        },
        emitEdit: function() {

            this.$emit('edit', this.todo);
        },
        emitDelete: function() {

            this.$emit('delete', {todo: this.todo, index: this.index});
        }
    },
    template: '<div class="todo-list-item level is-mobile">                                                             \
                <div class="level-left">                                                                                \
                    <div class="level-item">                                                                            \
                        <label class="checkbox">                                                                        \
                            <input type="checkbox"                                                                      \
                                    v-bind:checked="todo.done"                                                          \
                                    v-on:click="emitToggle">                                                            \
                        </label>                                                                                        \
                    </div>                                                                                              \
                    <div class="level-item">                                                                            \
                        <input type="text" class="input is-medium is-static" placeholder="Todo"                         \                                                                   \
                            v-on:click="$event.target.focus()"                                                          \
                            v-bind:class="{\'line-through\': todo.done}"                                                \
                            v-on:blur="emitEdit"                                                                        \
                            v-on:keydown.enter="emitEdit(); document.getElementById(\'note-input\').focus();"           \
                            v-model.trim="todo.note">                                                                   \
                    </div>                                                                                              \
                </div>                                                                                                  \
                <div class="level-right">                                                                               \
                    <div class="level-item">                                                                            \
                        <button class="button is-danger is-small"                                                       \
                                v-on:click="emitDelete">                                                                \
                                <span class="icon"><i class="mdi mdi-18px mdi-close"></i></span>                        \
                        </button>                                                                                       \
                    </div>                                                                                              \
                </div>                                                                                                  \
            </div>                                                                                                      \
    '
});



var app = new Vue({
  el: '#app',  
  data: {
    todos: [],
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
  methods: {
    orderChanged: function(order) {

        this.sortTodos();
        document.getElementById('note-input').focus();
    },
    sortTodos: function() {

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
    addTodo: function(text) {
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
    toggleTodo: function(todo) {

        todo.done = !todo.done;

        this.editTodo(todo);
    },
    editTodo: function(todo) {

        var xhttp = new XMLHttpRequest();
        xhttp.open("POST", "/edit/todo/"+todo.id, true);
        xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhttp.send(this.encodeTodo(todo));
        this.sortTodos();
    },
    deleteTodo: function(event) {

        var todo = event.todo;
        var index = event.index;
        this.todos.splice(index, 1);

        var xhttp = new XMLHttpRequest();
        xhttp.open("POST", "/delete/todo/"+todo.id, true);
        xhttp.send();
        document.getElementById('note-input').focus();
    }, 
    encodeTodo: function(todo) {

        return "note="+todo.note+"&done="+todo.done;
    }
  }
});