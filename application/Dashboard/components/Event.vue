<template>
    <span class="is-clickable" v-on:click="toggleExtended()">
        [{{ parseTimeStamp(event.timestamp) }}]: 
        <span v-bind:class="{ 'is-blue': isShopEvent, 'is-red': isCustomerEvent }">
            {{ event.code }} - {{ body }}
        </span>
        <pre v-if="extended">{{ event }}</pre>
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
    data: function()
    {
        return {
            extended: false
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

                return `${ this.event.body.name } had £${ this.event.body.money } and spent £${ this.event.body.spent } on ${ this.event.body.cart.length } products`
            }

            if(this.event.code === 'SHOP_SUPPLIED') 
            {
                let totalPrice = 0
                this.event.body.forEach((supply)=> {
                    return totalPrice += supply.price
                })
                return `Shop received ${ this.event.body.length } products for a total cost of £${ this.parseMoney(totalPrice) }`
            }

            if(this.event.code === 'SHOP_SOLD') 
            {
                return `Shop sold ${ this.event.body.length } products for £${ this.event.body.length }`
            }
        }
    },
    methods: {
        parseTimeStamp: function(value)
        {
            if(value)
            {
                return moment(String(value)).format('HH:mm')
            }
        },
        parseMoney: function(total)
        {
            if(total > 1000)
            {
                return `${ this.roundAmount(total / 1000) }k`
            }

            if(total > 1000000)
            {
                return `${ this.roundAmount(total / 1000000) }m`
            }

            return total
        },
        roundAmount: function(value)
        {
            return Math.round(value * 100) / 100
        },
        toggleExtended: function()
        {
            this.extended = !this.extended
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
.is-clickable {
    cursor: pointer;
}
</style>