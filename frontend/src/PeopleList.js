import React from 'react';
import Table from 'react-bootstrap/Table'
import axios from 'axios';

class PeopleList extends React.Component {

    readData() {
        const self = this;
        axios.get(window.global.api_location+'/people').then(function(response) {
            console.log(response.data);

            self.setState({people: response.data});
        }).catch(function (error){
            console.log(error);
        });
    }

    getPeople() {
        let table = []

        for (let i=0; i < this.state.people.length; i++) {

            table.push(
            <tr key={i}>
                <td>{this.state.people[i].firstName} {this.state.people[i].lastName}</td>
                <td>{this.state.people[i].email}</td>
                <td>{this.state.people[i].age}</td>
            </tr>
            );
        }

        return table
    }

    constructor(props) {
        super(props);
        this.readData();
        this.state = {people: []};
    
        this.readData = this.readData.bind(this);
    }

    render() {
      return (
        <div>
            <h1 style={{marginBottom: "40px"}}>Menu</h1>
            <Table>
                <thead>
                    <tr>
                        <th>
                            Name
                        </th>
                        <th>
                            Email
                        </th>
                        <th>
                            Age
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {this.getPeople()}
                </tbody>
            </Table>
        </div>
      ) 
    }
}

export default PeopleList;