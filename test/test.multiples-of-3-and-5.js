const { expect} = require('chai')
const { multiplesOf3And5 } = require('..')

describe('1: multiplesOf3And5()', () => {
  it('should find the sum of all the multiples of 3 or 5 below 1000.', () => {
    expect(multiplesOf3And5([3, 5], 1000)).to.equal(233168)
  })
})
