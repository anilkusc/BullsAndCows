
import React from 'react';
import Container from '@material-ui/core/Container';
import GameInfo from '../Components/GameInfo';
import MoveTable from '../Components/MoveTable';



class Game extends React.Component {
    render() {
        return (
            <div>
                <Container component="main">
                            <GameInfo/>
                </Container>
                <br></br>
                <br></br>
                <br></br>
                <Container component="main">
                            <MoveTable/>
                </Container>
            </div>
        );
    }
}
export default Game