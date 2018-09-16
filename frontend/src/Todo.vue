<template>
	<div class="todo-list-item level is-mobile">                                                             
	    <div class="level-left">                                                                                
	        <div class="level-item">                                                                            
	            <label class="checkbox">                                                                        
	                <input type="checkbox"                                                                      
	                        v-bind:checked="todo.done"                                                          
	                        v-on:click="emitToggle">                                                            
	            </label>                                                                                        
	        </div>                                                                                              
	        <div class="level-item">                                                                            
	            <input type="text" class="input is-medium is-static" placeholder="Todo"                                                                                            
	                v-on:click="$event.target.focus()"                                                          
	                v-bind:class="{'line-through': todo.done}"                                                
	                v-on:blur="emitEdit"                                                                        
	                v-on:keydown.enter="emitEdit(); document.getElementById('note-input').focus();"           
	                v-model.trim="todo.note">                                                                   
	        </div>                                                                                              
	    </div>                                                                                                  
	    <div class="level-right">                                                                               
	        <div class="level-item">                                                                            
	            <button class="button is-danger is-small"                                                       
	                    v-on:click="emitDelete">                                                                
	                    <span class="icon"><i class="mdi mdi-18px mdi-close"></i></span>                        
	            </button>                                                                                       
	        </div>                                                                                              
	    </div>                                                                                                  
	</div>	
</template>

<script>
export default {
	props: ["todo", "index"],
    methods: {
        emitToggle () {

            this.$emit('toggle', this.todo);
        },
        emitEdit () {

            this.$emit('edit', this.todo);
        },
        emitDelete () {

            this.$emit('delete', {todo: this.todo, index: this.index});
        }
    }
}
</script>