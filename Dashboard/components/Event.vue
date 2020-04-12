<template>
    <span v-bind:class="{ 'is-blue': isShopEvent, 'is-red': isCustomerEvent }">
        [{{ parseTimeStamp(event.timestamp) }}]: {{ event.code }} - {{ body }}
    </span>
</template>

<script>
import moment from 'moment'

export default {
    props: {
        event: {
            required: true,
            type: Object
        }
    },
    computed: {
        isShopEvent: function()
        {
            return this.event.service === 'shop'
        },
        isCustomerEvent: function()
        {
            return this.event.service === 'customer'
        },
        body: function()
        {
            if(this.event.code === 'CUSTOMER_BOUGHT') 
            {
                return `${ this.event.body.customer_id } had £${ this.event.body.money } and spent £${ this.event.body.spent } on ${ this.event.body.cart.length } products`
            }

            if(this.event.code === 'SHOP_SUPPLIED') 
            {
                return `Shop now have ${ this.event.body.quantity } of ${ this.event.body.product_id }`
            }

            if(this.event.code === 'SHOP_SOLD') 
            {
                return `Shop sold ${ this.event.body.cart.length } products for ${ this.event.body.spent }`
            }
        }
    },
    methods: {
        parseTimeStamp: function(value)
        {
            if(value)
            {
                return moment(String(value)).format('hh:mm:ss')
            }
        }
    }
}
</script>

<style scoped>
.is-blue {
    color: blue;
}
.is-red {
    color: red;
}
</style>