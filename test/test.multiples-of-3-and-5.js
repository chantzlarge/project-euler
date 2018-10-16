const { expect} = require('chai')
const { multiplesOf3And5 } = require('..')

describe('Multiples of 3 and 5', () => {
  it('Find the sum of all the multiples of 3 or 5 below 1000.', () => {
    expect(multiplesOf3And5([3, 5], 1000)).to.equal(233168)
  })
})
