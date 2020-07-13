import QtQuick 2.14
import QtQuick.Window 2.12
import QtQuik.Controls 2.12
ApplicationWindow
{
    id: root
    height: 200
    width: 400
    Rectangle
    {
        Button
        {
            property bool active: false
            id: startbutton
            anchors.left: parent.left 
            anchors.leftMargin: 10 
            text: "start"           
            onClicked: 
            {
                if  (active)
                {
                    active = false
                    return
                }


                
            }

        }
        Button
        {
            id: stopbutton
            anchors.left: parent.left 
            anchors.leftMargin: 10 
            text: "stop"
            

        }
  
    }
    Rectangle
    {
        height: 5
        length: root.width
    }
}
