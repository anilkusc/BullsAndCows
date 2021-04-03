import React from 'react';
import { DataGrid } from '@material-ui/data-grid';



class MoveTable extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            rows: [],
            columns: [
                { field: 'id', headerName: 'ID', width: 70 },
                { field: 'negative', headerName: 'Negative', width: 120 },
                { field: 'positive', headerName: 'Positive', width: 120 },
                { field: 'prediction', headerName: 'Prediction', width: 120 },
                { field: 'predictor', headerName: 'Predictor', width: 120 },
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