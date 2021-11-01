describe('Dopi Quickstart', () => {

    let text;

    beforeEach(() => {
        const rd = Math.floor(Math.random() * 1000);
        text = `Hello World ${rd}`
    })

    it('Should open item list', () => {
        cy.visit('http://localhost:8080/quickstart')
        cy.get('.item-list')
    })

    it('Should create item', () => {

        cy.visit('http://localhost:8080/quickstart')
        cy.get('#create-item-button').click()
        cy.get('.new-item')

        cy.get('#text').clear().type(text)
        cy.get('#save').click()

        cy.get('.item-list')
        cy.get('table').contains('td', text);
    })


})



