<template>
	<div class="level-item">                                                             
	    <span style="padding-right: 10px;">Order by: </span>                            
	    <div class="select"                                                             
	            v-bind:class="{'is-loading': loader}">                                
	        <select v-model="value" v-on:input="selectChanged($event.target.value)">    
	            <option value="created_at">Date</option>                                
	            <option value="note">Name</option>                                      
	        </select>                                                                   
	    </div>                                                                          
	</div>	
</template>

<script>
export default {
	props: {
        loader: Boolean,
        value: {
            type: String,
            default: "created_at"
        }
    },
    methods: {
        selectChanged (order) {
            this.$emit('input', order);
            this.$emit('order-changed', order);
        }
    }
}	
</script>