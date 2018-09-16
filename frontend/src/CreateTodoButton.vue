<template>
	<div class="todo-list-item level is-mobile">                                                         
	    <div class="level-left">                                                                            
	        <div class="level-item">                                                                        
	            <label class="checkbox"><input type="checkbox" disabled></label>                            
	        </div>                                                                                          
	        <div class="leve-item">                                                                         
	            <input id="note-input" class="input is-medium is-static" type="text" placeholder="New Todo" 
	                v-model.trim="text"                                                                     
	                v-on:keydown.enter="emitCreateTodo">                                                    
	        </div>                                                                                          
	    </div>                                                                                              
	    <div class="level-right">                                                                           
	        <div class="level-item">                                                                        
	            <button class="button is-success is-medium" v-on:click="emitCreateTodo">                    
	                <span class="icon">                                                                     
	                    <i class="mdi mdi-24px mdi-plus"></i>                                               
	                </span>                                                                                 
	                <span>Create</span>                                                                     
	            </button>                                                                                   
	        </div>                                                                                          
	    </div>                                                                                              
	</div> 	
</template>

<script>
export default {
	data () {
        return {
            text: "",
        };
    },
    methods: {
        emitCreateTodo () {
            if(this.text != "") {

                this.$emit('create-todo', this.text);
                this.text = "";
            }
        }
    }
}
</script>