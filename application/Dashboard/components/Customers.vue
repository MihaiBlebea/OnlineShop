<template>
    <div class="card">
        <div class="card-body">
            <p class="d-flex justify-content-between">
                <span>Total</span>
                <span>{{ totalCount }} customers</span>
            </p>
            <p class="d-flex justify-content-between">
                <span>Total spent</span>
                <span>£{{ totalSpent }}</span>
            </p>
            <p class="d-flex justify-content-between mb-0">
                <span>Avr. customer spent</span>
                <span>£{{ avrCustomerSpent }}</span>
            </p>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        data: {
            type: Array,
            required: false,
            default: []
        }
    },
    data: function()
    {
        return {
            spend: 0
        }
    },
    computed: {
        totalCount: function()
        {
            console.log('AICI', this.data)
            if(!this.data || this.data.length === 0)
            {
                return 0
            }
            return this.data.length
        },
        totalSpent: function() 
        {
            if(!this.data || this.data.length === 0)
            {
                return 0
            }
            let total = 0
            this.data.forEach((customer)=> {
                total += customer.spent
            })
            this.spend = total
            if(total > 1000)
            {
                return `${ this.roundAmount(total / 1000) }k`
            }

            if(total > 1000000)
            {
                return `${ this.roundAmount(total / 1000000) }m`
            }

            return this.roundAmount(total)
        },
        avrCustomerSpent: function()
        {
            if(!this.data || this.data.length === 0 || this.spend === 0)
            {
                return 0
            }
            return this.roundAmount(this.spend / this.data.length)
        }
    },
    methods: 
    {
        roundAmount: function(value)
        {
            return Math.round(value * 100) / 100
        }
    }
}
</script>