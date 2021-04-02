import React from 'react';
import CreateGame from '../Components/CreateGame'
import JoinGame from '../Components/JoinGame'
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';

class Main extends React.Component {
    render() {
        return (
            <div>
                <Container component="main">
                    <Grid container justify="center" spacing={12} >
                        <Grid item md={6}>
                            <CreateGame />
                        </Grid>
                        <Grid item md={6}>
                            <JoinGame />
                        </Grid>
                    </Grid>
                </Container>


            </div>
        );
    }
}
export default Main