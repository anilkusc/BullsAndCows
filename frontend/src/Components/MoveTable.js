import React from 'react';
import { DataGrid } from '@material-ui/data-grid';



class MoveTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            rows: [],
            columns: [
                { field: 'id', headerName: 'ID', width: 70 },
                { field: 'username', headerName: 'Username', width: 130 },
                { field: 'role', headerName: 'Role', width: 130 },
            ]
        }
    }

    render() {
        return (

                <div style={{ height: 400, width: '100%' }}>
                    <DataGrid rows={this.state.rows} columns={this.state.columns} pageSize={5} />
                </div>
        );
    }
}
export default (MoveTable)