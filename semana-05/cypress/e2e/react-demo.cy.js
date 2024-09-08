//react-demo.cy.js
describe('Wikipedia text example', () => {
  it('Should validate h1 text', () => {
  cy.visit('https://pt.wikipedia.org/wiki/Jogos_Paral%C3%ADmpicos_de_Ver%C3%A3o_de_2024');
  cy.get('h1 > span').should('have.text','Jogos Paralímpicos de Verão de 2024')
  })
})