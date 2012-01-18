#include "qtsignal.h"
#include "cdrv.h"
#include <QDebug>

QtSignal::QtSignal(QObject *parent, void *func) :
    QObject(parent), m_func(func)
{
}

QtSignal::~QtSignal()
{
}

void QtSignal::call()
{
    drv_callback(m_func,0,0,0,0);
}

void QtSignal::call(bool b)
{
    int i = b ? 1:0;
    drv_callback(m_func,&i,0,0,0);
}

void QtSignal::call(int i)
{
    drv_callback(m_func,&i,0,0,0);
}

void QtSignal::call(const QString& s)
{
    string_head sh1;
    drvSetString(&sh1,s);
    drv_callback(m_func,&sh1,0,0,0);
}

void QtSignal::call(QObject*)
{
    //((void (*)(void*))(m_func))(m_param);
}
