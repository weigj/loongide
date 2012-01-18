#ifndef QTAPP_H
#define QTAPP_H

#include <QApplication>

class QtApp : public QApplication
{
    Q_OBJECT
public:
    explicit QtApp(int argc = 0, char **argv = NULL);
        
public slots:
    void deleteObject(QObject *obj);
public:
    void *pfnDeleteObj;
};

#endif // QTAPP_H
