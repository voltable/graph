import React from 'react';
import { render } from 'react-dom';
import RaisedButton from 'material-ui/RaisedButton';
import Dialog from 'material-ui/Dialog';
import { deepOrange500 } from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import D3visualization from './modules/D3Visualization';

const styles = {
    container: {
        textAlign: 'center',
        paddingTop: 200,
    },
};

const muiTheme = getMuiTheme({
    palette: {
        accent1Color: deepOrange500,
    },
});

const App =
    <MuiThemeProvider muiTheme={muiTheme}>
        <div style={styles.container}>
            <Dialog

                title="Super Secret Password"
            >
                1-2-3-4-5
           
            </Dialog>
            <D3visualization></D3visualization>
            <h1>Material-UI</h1>
            <h2>example project</h2>
            <RaisedButton
                label="Super Secret Password"
                secondary={true}
            />
        </div>
    </MuiThemeProvider>;


export default App