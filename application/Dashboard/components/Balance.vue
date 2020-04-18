<template>
    <div class="card">
        <div class="card-body">
            <p class="d-flex justify-content-between mb-0">
                <span>Balance</span>
                <span v-bind:class="{ 'is-positive': isPositive, 'is-negative': !isPositive }">£{{ parseMoney(balance) }}</span>
            </p>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        balance: {
            required: true,
            type: Number,
            default: 0
        },
        transactions: {
            required: false,
            default: [],
            type: Array
        }
    },
    computed: {
        isPositive: function()
        {
            return this.balance > 0
        }
    },
    methods: {
        parseMoney: function(total)
        {
            if(Math.abs(total) > 1000)
            {
                return `${ this.roundAmount(total / 1000) }k`
            }

            if(Math.abs(total) > 1000000)
            {
                return `${ this.roundAmount(total / 1000000) }m`
            }

            return total
        },
        roundAmount: function(value)
        {
            return Math.round(value * 100) / 100
        },
    }
}
</script>

<style scoped>
.is-positive {
    color: green;
}
.is-negative {
    color: red;
}
</style>>