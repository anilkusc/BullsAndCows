import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';

const useStyles = theme => ({
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    form: {
        width: '100%', // Fix IE 11 issue.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },

});

class GameInfo extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            start: false,
        }
    }

    render() {
        const { classes } = this.props;
        return (
            <div>
                <Container component="main" maxWidth="xs">
                    <div className={classes.paper}>
                        <Paper>
                            &nbsp;&nbsp;&nbsp;
                            Turn:
                            &nbsp;&nbsp;&nbsp;
                            Move:
                            &nbsp;&nbsp;&nbsp;
                            Opponent:
                            &nbsp;&nbsp;&nbsp;
                            Session:
                            &nbsp;&nbsp;&nbsp;
                            <Button>Abandon</Button>

                            <Container component="main" maxWidth="xs">
                                {this.state.start ? (
                                    <div className={classes.paper}>
                                        <Typography component="h1" variant="h5">
                                            Your Prediction
                                        </Typography>
                                        <form className={classes.form} noValidate onSubmit={this.handleSubmit}>
                                            <TextField
                                                variant="outlined"
                                                margin="normal"
                                                color="primary"
                                                required
                                                fullWidth
                                                id="text"
                                                label="Prediction"
                                                name="prediction"
                                                autoComplete="prediction"
                                                onChange={this.handleChangeUsername}
                                                autoFocus
                                            />
                                            <Button
                                                type="submit"
                                                fullWidth
                                                variant="contained"
                                                color="primary"
                                                onClick={this.handleSubmit}
                                                className={classes.submit}
                                            >
                                                SEND
                                            </Button>
                                        </form>
                                    </div>
                                ) :
                                    (
                                        <Typography component="h1" variant="h5">
                                            <br></br>
                                            <br></br>
                                            Waiting For The Opponent...
                                            <br></br>
                                            <br></br>
                                            <br></br>
                                        </Typography>
                                    )
                                }
                            </Container>
                        </Paper>
                    </div>
                </Container>
            </div>
        );
    }
}
export default withStyles(useStyles)(GameInfo)