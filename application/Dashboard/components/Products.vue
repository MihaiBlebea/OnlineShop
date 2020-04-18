<template>
    <div class="card">
        <div class="card-body">
            <p class="d-flex justify-content-between">
                <span>Uniq. in stock</span>
                <span>{{ uniqCount }}</span>
            </p>
            <p class="d-flex justify-content-between">
                <span>Total # of products</span>
                <span>{{ totalCount }}</span>
            </p>
            <p class="d-flex justify-content-between">
                <span>Avr. product value</span>
                <span>£{{ avrValue }}</span>
            </p>
            <p class="d-flex justify-content-between mb-0">
                <span>£ total value</span>
                <span>£{{ totalValue }}</span>
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
        uniqCount: function()
        {
            if(!this.data || this.data.length === 0)
            {
                return 0
            }
            return this.data.length
        },
        totalValue: function() 
        {
            if(!this.data || this.data.length === 0)
            {
                return 0
            }
            let total = 0
            this.data.forEach((product)=> {
                total += product.quantity * parseFloat(product.price)
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
        totalCount: function()
        {
            if(!this.data || this.data.length === 0)
            {
                return 0
            }
            let count = 0
            this.data.forEach((product)=> {
                count += product.quantity
            })
            return count
        },
        avrValue: function()
        {
            if(!this.data || this.data.length === 0 || this.totalCount === 0)
            {
                return 0
            }
            return this.roundAmount(this.spend / this.totalCount)
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