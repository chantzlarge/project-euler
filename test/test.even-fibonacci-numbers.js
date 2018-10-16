const { expect} = require('chai')
const { evenFibonacciNumbers } = require('..')

describe('Even Fibonacci Numbers', () => {
  it('By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.', () => {
    expect(evenFibonacciNumbers([1, 2], 4000000)).to.equal(4613732)
  })
})
