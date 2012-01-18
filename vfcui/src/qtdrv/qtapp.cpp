#include "qtapp.h"
#include "cdrv.h"

QtApp::QtApp(int argc, char **argv) :
    QApplication(argc,argv), pfnDeleteObj(0)
{
}

void QtApp::deleteObject(QObject *obj)
{
    drv_callback(pfnDeleteObj,obj,0,0,0);
}
